# VERSION defines the project version for the bundle.
# Update this value when you upgrade the version of your project.
# To re-generate a bundle for another specific version without changing the standard setup, you can:
# - use the VERSION as arg of the bundle target (e.g make bundle VERSION=0.0.2)
# - use environment variables to overwrite this value (e.g export VERSION=0.0.2)
VERSION ?= $(shell git branch --show-current)
IMG_VERSION ?= $(shell git branch --show-current)
OPERATOR_VER := $(shell git branch --show-current)

EXTRA_SERVICE_ACCOUNTS := --extra-service-accounts="sts-plugin,sts-tsync"

TSYNC_VERSION := 2.1.2.9
ICE_VERSION = 1.8.8

MARKETPLACE_REMOTE_WORKFLOW  := https://marketplace.redhat.com/en-us/operators/silicom-sts-operator/pricing?utm_source=openshift_console
MARKETPLACE_SUPPORT_WORKFLOW := https://marketplace.redhat.com/en-us/operators/silicom-sts-operator/support?utm_source=openshift_console

PREFLIGHT_TARGETS := preflight-tsyncd
PREFLIGHT_TARGETS += preflight-operator
PREFLIGHT_TARGETS += preflight-plugin
PREFLIGHT_TARGETS += preflight-phc2sys
PREFLIGHT_TARGETS += preflight-gpsd
PREFLIGHT_TARGETS += preflight-grpc-tsyncd

.PHONY: $(PREFLIGHT_TARGETS)

export PFLT_LOGLEVEL=debug
export PFLT_INDEXIMAGE=$(IMAGE_REGISTRY)/sts-operator-catalog:$(IMG_VERSION)
export PFLT_DOCKERCONFIG=$(shell pwd)/docker.json

# CHANNELS define the bundle channels used in the bundle.
# Add a new line here if you would like to change its default config. (E.g CHANNELS = "candidate,fast,stable")
# To re-generate a bundle for other specific channels without changing the standard setup, you can:
# - use the CHANNELS as arg of the bundle target (e.g make bundle CHANNELS=candidate,fast,stable)
# - use environment variables to overwrite this value (e.g export CHANNELS="candidate,fast,stable")
CHANNELS?="alpha,stable"
ifneq ($(origin CHANNELS), undefined)
BUNDLE_CHANNELS := --channels=$(CHANNELS)
endif

# DEFAULT_CHANNEL defines the default channel used in the bundle.
# Add a new line here if you would like to change its default config. (E.g DEFAULT_CHANNEL = "stable")
# To re-generate a bundle for any other default channel without changing the default setup, you can:
# - use the DEFAULT_CHANNEL as arg of the bundle target (e.g make bundle DEFAULT_CHANNEL=stable)
# - use environment variables to overwrite this value (e.g export DEFAULT_CHANNEL="stable")
DEFAULT_CHANNEL?="stable"
ifneq ($(origin DEFAULT_CHANNEL), undefined)
BUNDLE_DEFAULT_CHANNEL := --default-channel=$(DEFAULT_CHANNEL)
endif
BUNDLE_METADATA_OPTS ?= $(BUNDLE_CHANNELS) $(BUNDLE_DEFAULT_CHANNEL)

# IMAGE_TAG_BASE defines the docker.io namespace and part of the image name for remote images.
# This variable is used to construct full image tags for bundle and catalog images.
#
# For example, running 'make bundle-build bundle-push catalog-build catalog-push' will build and push both
# $(IMAGE_REGISTRY)/operator-bundle:$VERSION and quay.io/silicom/operator-catalog:$VERSION.
IMAGE_REGISTRY ?= quay.io/silicom
IMAGE_TAG_BASE ?= $(IMAGE_REGISTRY)/sts-operator

# BUNDLE_IMG defines the image:tag used for the bundle.
# You can use it as an arg. (E.g make bundle-build BUNDLE_IMG=<some-registry>/<project-name-bundle>:<tag>)
BUNDLE_IMG ?= $(IMAGE_TAG_BASE)-bundle:$(VERSION)

# Image URL to use all building/pushing image targets
IMG ?= $(IMAGE_TAG_BASE):$(IMG_VERSION)
# Produce CRDs that work back to Kubernetes 1.11 (no version conversion)
CRD_OPTIONS ?= "crd:trivialVersions=true,preserveUnknownFields=false"
# ENVTEST_K8S_VERSION refers to the version of kubebuilder assets to be downloaded by envtest binary.
ENVTEST_K8S_VERSION = 1.21

COMMUNITY_PROD_OPERATORS_GIT := https://github.com/rmr/community-operators-prod.git
COMMUNITY_PROD_OPERATORS_DIR := community-operators-prod
OPERATOR_NAME			:= silicom-sts-operator

CERTIFIED_DIR         := certified-operators
CERTIFIED_GIT         := https://github.com/rmr/certified-operators.git

# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif

# Setting SHELL to bash allows bash commands to be executed by recipes.
# This is a requirement for 'setup-envtest.sh' in the test target.
# Options are set to exit when a recipe line exits non-zero or a piped command fails.
SHELL = /usr/bin/env bash -o pipefail
.SHELLFLAGS = -ec

all: build controller-gen $(PREFLIGHT)

$(PREFLIGHT_TARGETS): $(PREFLIGHT)

##@ General

# The help target prints out all targets with their descriptions organized
# beneath their categories. The categories are represented by '##@' and the
# target descriptions by '##'. The awk commands is responsible for reading the
# entire set of makefiles included in this invocation, looking for lines of the
# file as xyz: ## something, and then pretty-format the target and help. Then,
# if there's a line with ##@ something, that gets pretty-printed as a category.
# More info on the usage of ANSI control characters for terminal formatting:
# https://en.wikipedia.org/wiki/ANSI_escape_code#SGR_parameters
# More info on the awk command:
# http://linuxcommand.org/lc3_adv_awk.php

help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Development

manifests: controller-gen ## Generate WebhookConfiguration, ClusterRole and CustomResourceDefinition objects.
	$(CONTROLLER_GEN) $(CRD_OPTIONS) rbac:roleName=manager-role webhook paths="./..." output:crd:artifacts:config=config/crd/bases

generate: controller-gen ## Generate code containing DeepCopy, DeepCopyInto, and DeepCopyObject method implementations.
	$(CONTROLLER_GEN) object:headerFile="hack/boilerplate.go.txt" paths="./..."

fmt: ## Run go fmt against code.
	go fmt ./...

vet: ## Run go vet against code.
	go vet ./...

test: manifests generate fmt vet envtest ## Run tests.
	KUBEBUILDER_ASSETS="$(shell $(ENVTEST) use $(ENVTEST_K8S_VERSION) -p path)" go test ./... -coverprofile cover.out

test-kind:
	make uninstall
	make install
	bash test/run-kind.sh

##@ Build

.PHONY: build
build: controller-gen operator-sdk preflight generate fmt vet ## Build manager binary.
	go build -o bin/manager main.go

run: manifests generate fmt vet ## Run a controller from your host.
	go run ./main.go --zap-devel

docker-build: test ## Build docker image with the manager.
	docker build --build-arg STS_VERSION=$(VERSION) -t ${IMG} .

docker-push: ## Push docker image with the manager.
	docker push ${IMG}

##@ Deployment

YQ := bin/yq
yq: # Install yq if necessary to parse and manipulate yaml files.
	curl -sL https://github.com/mikefarah/yq/releases/download/v4.24.5/yq_linux_amd64 -o $(YQ)
	chmod +x $(YQ)

ACT := bin/act
act: # Install act if necessary to locally test github CI/CD.
	curl https://raw.githubusercontent.com/nektos/act/master/install.sh | bash

install: manifests kustomize ## Install CRDs into the K8s cluster specified in ~/.kube/config.
	$(KUSTOMIZE) build config/crd | kubectl apply -f -

uninstall: manifests kustomize ## Uninstall CRDs from the K8s cluster specified in ~/.kube/config.
	$(KUSTOMIZE) build config/crd | kubectl delete -f -

deploy: manifests kustomize ## Deploy controller to the K8s cluster specified in ~/.kube/config.
	cd config/manager && $(KUSTOMIZE) edit set image controller=${IMG}
	$(KUSTOMIZE) build config/default | kubectl apply -f -

undeploy: ## Undeploy controller from the K8s cluster specified in ~/.kube/config.
	$(KUSTOMIZE) build config/default | kubectl delete -f -

PREFLIGHT = bin/preflight
preflight: controller-gen kustomize  ## Download opm locally if necessary.
	curl -sL https://github.com/redhat-openshift-ecosystem/openshift-preflight/releases/download/1.1.0/preflight-linux-amd64 -o ./bin/preflight
	chmod +x bin/preflight

.PHONY: controller-gen
CONTROLLER_GEN = bin/controller-gen
controller-gen: ## Download controller-gen locally if necessary.
	$(call go-get-tool,$(CONTROLLER_GEN),sigs.k8s.io/controller-tools/cmd/controller-gen@v0.6.1)

KUSTOMIZE = $(shell pwd)/bin/kustomize
kustomize: ## Download kustomize locally if necessary.
	$(call go-get-tool,$(KUSTOMIZE),sigs.k8s.io/kustomize/kustomize/v3@v3.8.7)

ENVTEST = $(shell pwd)/bin/setup-envtest
envtest: ## Download envtest-setup locally if necessary.
	$(call go-get-tool,$(ENVTEST),sigs.k8s.io/controller-runtime/tools/setup-envtest@latest)

# go-get-tool will 'go get' any package $2 and install it to $1.
PROJECT_DIR := $(shell dirname $(abspath $(lastword $(MAKEFILE_LIST))))
define go-get-tool
@[ -f $(1) ] || { \
set -ex ;\
TMP_DIR=$$(mktemp -d) ;\
cd $$TMP_DIR ;\
go mod init tmp ;\
echo "Downloading $(2)" ;\
GOBIN=$(PROJECT_DIR)/bin go get $(2) ;\
rm -rf $$TMP_DIR ;\
}
endef

.PHONY: operator-sdk
OPERATOR_SDK = $(shell pwd)/bin/operator-sdk
operator-sdk: bin
	curl -sL https://github.com/operator-framework/operator-sdk/releases/download/v1.19.1/operator-sdk_linux_amd64 -o bin/operator-sdk
	chmod +x bin/operator-sdk

.PHONY: all bundle
bundle: manifests kustomize ## Generate bundle manifests and metadata, then validate generated files.
#	- rm  bundle/manifests/silicom-sts-operator.clusterserviceversion.yaml
	sed -i 's/sts-plugin:.*[^"]/sts-plugin:$(OPERATOR_VER)/g' api/v1alpha1/stsoperatorconfig_types.go
	sed -i 's/tsyncd:.*[^"}]/tsyncd:$(TSYNC_VERSION)/g' api/v1alpha1/stsoperatorconfig_types.go
	sed -i 's/grpc-tsyncd:.*[^"}]/grpc-tsyncd:$(TSYNC_VERSION)/g' api/v1alpha1/stsoperatorconfig_types.go
	bin/operator-sdk generate kustomize manifests -q
	cd config/manager && $(KUSTOMIZE) edit set image controller=$(IMG)
	$(KUSTOMIZE) build config/manifests | bin/operator-sdk generate bundle --overwrite -q --version $(VERSION) $(BUNDLE_METADATA_OPTS) $(EXTRA_SERVICE_ACCOUNTS)
	echo "  com.redhat.openshift.versions: \"v4.8\"" >> bundle/metadata/annotations.yaml
	echo "LABEL com.redhat.openshift.versions=\"v4.8\"" >> bundle.Dockerfile
	echo "LABEL com.redhat.delivery.operator.bundle=true" >> bundle.Dockerfile
	cat images.yaml >> bundle/manifests/silicom-sts-operator.clusterserviceversion.yaml
	rm bundle/manifests/*-config_v1_configmap.yaml
	bin/operator-sdk bundle validate ./bundle

.PHONY: bundle-build
bundle-build: ## Build the bundle image.
	docker build -f bundle.Dockerfile -t $(BUNDLE_IMG) .

.PHONY: bundle-push
bundle-push: ## Push the bundle image.
	$(MAKE) docker-push IMG=$(BUNDLE_IMG)

bin:
	mkdir bin

.PHONY: opm
OPM = ./bin/opm
opm: ## Download opm locally if necessary.
ifeq (,$(wildcard $(OPM)))
ifeq (,$(shell which opm 2>/dev/null))
	@{ \
	set -e ;\
	mkdir -p $(dir $(OPM)) ;\
	OS=$(shell go env GOOS) && ARCH=$(shell go env GOARCH) && \
	curl -sSLo $(OPM) https://github.com/operator-framework/operator-registry/releases/download/v1.15.1/$${OS}-$${ARCH}-opm ;\
	chmod +x $(OPM) ;\
	}
else
OPM = $(shell which opm)
endif
endif

# A comma-separated list of bundle images (e.g. make catalog-build BUNDLE_IMGS=example.com/operator-bundle:v0.1.0,example.com/operator-bundle:v0.2.0).
# These images MUST exist in a registry and be pull-able.
BUNDLE_IMGS ?= $(BUNDLE_IMG)

# The image tag given to the resulting catalog image (e.g. make catalog-build CATALOG_IMG=example.com/operator-catalog:v0.2.0).
CATALOG_IMG ?= $(IMAGE_TAG_BASE)-catalog:$(VERSION)

# Set CATALOG_BASE_IMG to an existing catalog image tag to add $BUNDLE_IMGS to that image.
ifneq ($(origin CATALOG_BASE_IMG), undefined)
FROM_INDEX_OPT := --from-index $(CATALOG_BASE_IMG)
endif

# Build a catalog image by adding bundle images to an empty catalog using the operator package manager tool, 'opm'.
# This recipe invokes 'opm' in 'semver' bundle add mode. For more information on add modes, see:
# https://github.com/operator-framework/community-operators/blob/7f1438c/docs/packaging-operator.md#updating-your-existing-operator
.PHONY: catalog-build
catalog-build: opm bundle ## Build a catalog image.
	$(OPM) index add --container-tool docker --mode semver --tag $(CATALOG_IMG) --bundles $(BUNDLE_IMGS) $(FROM_INDEX_OPT)

# Push the catalog image.
.PHONY: catalog-push
catalog-push: ## Push a catalog image.
	$(MAKE) docker-push IMG=$(CATALOG_IMG)

##@ Openshift preflight

.PHONY: preflight-all $(PREFLIGHT_TARGETS)
preflight-all: $(PREFLIGHT_TARGETS)
preflight-plugin: ## Check sts-plugin with preflight and upload results.
	$(PREFLIGHT) check container \
		$(shell $(YQ) '.relatedImages.[] | select(.name == "sts-plugin") | .image ' images.yaml) \
		--certification-project-id=627233b4d3d88855cdc646f5 \
		--docker-config=$(shell pwd)/config.json \
		--submit

preflight-gpsd: ## Check gpsd with preflight and upload results.
	$(PREFLIGHT) check container \
		$(shell $(YQ) '.relatedImages.[] | select(.name == "gpsd") | .image ' images.yaml) \
		--certification-project-id=622b5495f8469c36ac475618 \
		--docker-config=$(shell pwd)/config.json \
		--submit

preflight-tsyncd: ## Check tsyncd with preflight and upload results.
	$(PREFLIGHT) check container \
		$(shell $(YQ) '.relatedImages.[] | select(.name == "tsyncd") | .image ' images.yaml) \
		--certification-project-id=6218dc7622ee06da01c10bb5 \
		--docker-config=$(shell pwd)/config.json \
		--submit

preflight-grpc-tsyncd: ## Check grpc-tsyncd with preflight and upload results.
	$(PREFLIGHT) check container \
		$(shell $(YQ) '.relatedImages.[] | select(.name == "grpc-tsyncd") | .image ' images.yaml) \
		--certification-project-id=62651e90e6f5b76c831ba804 \
		--docker-config=$(shell pwd)/config.json \
		--submit

preflight-phc2sys: ## Check phc2sys with preflight and upload results.
	$(PREFLIGHT) check container \
		$(shell $(YQ) '.relatedImages.[] | select(.name == "phc2sys") | .image ' images.yaml) \
		--certification-project-id 6265110a59837e5a2f051c39 \
		--docker-config $(shell pwd)/config.json \
		--submit

preflight-operator-image: ## Check sts-operator with preflight and upload results.
	$(PREFLIGHT) check container \
		$(shell $(YQ) '.relatedImages.[] | select(.name == "sts-operator") | .image ' images.yaml) \
		--certification-project-id 6268270b61336b5931b96337 \
		--docker-config $(shell pwd)/config.json \
		--submit

preflight-operator-bundle: ## Check operator-bundle with preflight and upload results.
	which operator-sdk
	$(PREFLIGHT) check operator  \
		$(shell docker inspect $(IMAGE_REGISTRY)/sts-operator-bundle:$(IMG_VERSION) --format '{{ index .RepoDigests 0 }}')

preflight-ice-driver:
	$(PREFLIGHT) check container \
		$(shell docker inspect $(IMAGE_REGISTRY)/ice-driver-src:$(ICE_VERSION) --format '{{ index .RepoDigests 0 }}') \
		--certification-project-id=62669911d634ea6b75a3af8b \
		--docker-config=$(shell pwd)/config.json \
		--submit

plugin:
	docker build . -t $(IMAGE_REGISTRY)/sts-plugin:$(IMG_VERSION) \
		--build-arg STS_VERSION=$(VERSION) \
		-f Dockerfile.plugin

all-sts:
	rm -rf bundle
	make docker-build
	make docker-push
	make plugin
	make plugin-push
	make update-images
	make bundle
	make bundle-build
	make bundle-push
	make catalog-build
	make catalog-push
	make preflight-all
	-rm -rf certified-operators
	make certified-clone
	mkdir -p certified-operators/operators/silicom-sts-operator/$(IMG_VERSION)
	make certified-bundle

plugin-push:
	docker push $(IMAGE_REGISTRY)/sts-plugin:$(IMG_VERSION)

bundle-all: generate manifests bundle bundle-build

community-clone:
	git clone $(COMMUNITY_PROD_OPERATORS_GIT)

community-bundle: bundle
	cp bundle.Dockerfile  $(COMMUNITY_PROD_OPERATORS_DIR)/operators/$(OPERATOR_NAME)/$(OPERATOR_VER)/
	cp -av bundle/* $(COMMUNITY_PROD_OPERATORS_DIR)/operators/$(OPERATOR_NAME)/$(OPERATOR_VER)/


certified-clone:
	git clone $(CERTIFIED_GIT)

certified-bundle: bundle
	cp bundle.Dockerfile  $(CERTIFIED_DIR)/operators/$(OPERATOR_NAME)/$(OPERATOR_VER)/
	cp -av bundle/* $(CERTIFIED_DIR)/operators/$(OPERATOR_NAME)/$(OPERATOR_VER)/
	@echo "cert_project_id: 6266943761336b5931b9632c" > $(CERTIFIED_DIR)/operators/$(OPERATOR_NAME)/ci.yaml
	sed -i 's/sts-operator:$(OPERATOR_VER)/sts-operator@$(shell $(YQ) '.relatedImages.[] | select(.name == "sts-operator") | .image ' images.yaml | cut -d '@' -f 2)/' \
			 $(CERTIFIED_DIR)/operators/$(OPERATOR_NAME)/$(OPERATOR_VER)/manifests/silicom-sts-operator.clusterserviceversion.yaml
	sed -i '/replaces:/d' $(CERTIFIED_DIR)/operators/$(OPERATOR_NAME)/$(OPERATOR_VER)/manifests/silicom-sts-operator.clusterserviceversion.yaml

OPP = bin/opp.sh
opp:
	curl -sL https://raw.githubusercontent.com/redhat-openshift-ecosystem/community-operators-pipeline/ci/latest/ci/scripts/opp.sh -o bin/opp.sh
	chmod +x bin/opp.sh

opp-community-test: community-bundle
	cd $(COMMUNITY_OPERATORS_DIR)
	OPP_PRODUCTION_TYPE=ocp OPP_AUTO_PACKAGEMANIFEST_CLUSTER_VERSION_LABEL=1 \
		$(OPP) all $(COMMUNITY_OPERATORS_DIR)/operators/$(OPERATOR_NAME)/$(OPERATOR_VER)

update-images:
	@echo "$(shell docker pull -q $(IMAGE_REGISTRY)/gpsd:3.23.1)"
	@echo "$(shell docker pull -q $(IMAGE_REGISTRY)/sts-plugin:$(OPERATOR_VER))"
	@echo "$(shell docker pull -q $(IMAGE_REGISTRY)/tsyncd:$(TSYNC_VERSION) )"
	@echo "$(shell docker pull -q $(IMAGE_REGISTRY)/grpc-tsyncd:$(TSYNC_VERSION))"
	@echo "$(shell docker pull -q $(IMAGE_REGISTRY)/phc2sys:3.1.1)"
	@echo "$(shell docker pull -q $(IMAGE_REGISTRY)/ice-driver-src:$(ICE_VERSION))"
	@echo "$(shell docker pull -q $(IMAGE_REGISTRY)/sts-operator:$(OPERATOR_VER))"
	@echo "$(shell docker pull -q gcr.io/kubebuilder/kube-rbac-proxy:v0.8.0)"
	@echo "  relatedImages:" > images.yaml
	@echo "  - image: $(shell docker inspect $(IMAGE_REGISTRY)/gpsd:3.23.1 --format '{{ index .RepoDigests 0 }}')" >> images.yaml
	@echo "    name: gpsd" >> images.yaml
	@echo "  - image: $(shell docker inspect $(IMAGE_REGISTRY)/phc2sys:3.1.1 --format '{{ index .RepoDigests 0 }}')" >> images.yaml
	@echo "    name: phc2sys" >> images.yaml
	@echo "  - image: $(shell docker inspect $(IMAGE_REGISTRY)/tsyncd:$(TSYNC_VERSION)  --format '{{ index .RepoDigests 0 }}')" >> images.yaml
	@echo "    name: tsyncd" >> images.yaml
	@echo "  - image: $(shell docker inspect $(IMAGE_REGISTRY)/grpc-tsyncd:$(TSYNC_VERSION) --format '{{ index .RepoDigests 0 }}')" >> images.yaml
	@echo "    name: grpc-tsyncd" >> images.yaml
	@echo "  - image: $(shell docker inspect $(IMAGE_REGISTRY)/sts-plugin:$(VERSION) --format '{{ index .RepoDigests 0 }}')" >> images.yaml
	@echo "    name: sts-plugin" >> images.yaml
	@echo "  - image: $(shell docker inspect $(IMAGE_REGISTRY)/ice-driver-src:$(ICE_VERSION) --format '{{ index .RepoDigests 0 }}')" >> images.yaml
	@echo "    name: ice-driver-src" >> images.yaml
	@echo "  - image: $(shell docker inspect $(IMAGE_REGISTRY)/sts-operator:$(OPERATOR_VER) --format '{{ index .RepoDigests 0 }}')" >> images.yaml
	@echo "    name: sts-operator" >> images.yaml
	@echo "  - image: $(shell docker inspect gcr.io/kubebuilder/kube-rbac-proxy:v0.8.0 --format '{{ index .RepoDigests 0 }}')" >> images.yaml
	@echo "    name: kube-rbac-proxy" >> images.yaml
