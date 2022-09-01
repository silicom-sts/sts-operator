module github.com/rmr/sts-operator

go 1.16

require (
	github.com/Masterminds/sprig v2.20.0+incompatible
	github.com/go-logr/logr v0.4.0
	github.com/onsi/ginkgo v1.16.5
	github.com/onsi/gomega v1.15.0
	github.com/openshift-psap/special-resource-operator v0.0.0-20211207075758-ecc22857e107
	github.com/openshift/cluster-nfd-operator v0.0.0-20211111213742-99325d66ab18
	k8s.io/api v0.22.2
	k8s.io/apimachinery v0.22.2
	k8s.io/client-go v0.22.2
	k8s.io/utils v0.0.0-20210819203725-bdf08cb9a70a
	sigs.k8s.io/controller-runtime v0.10.2
)
