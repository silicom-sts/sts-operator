#!/bin/bash

scriptdir="$( dirname -- "${BASH_SOURCE[0]}"; )";

kubectl label nodes kind-control-plane sts.silicom.com/config=bc-1 --overwrite
kubectl apply -f https://raw.githubusercontent.com/openshift/cluster-nfd-operator/release-4.10/manifests/4.10/manifests/nfd.openshift.io_nodefeaturediscoveries.yaml
kubectl apply -f ${scriptdir}/rbac.yaml
kubectl apply -f ${scriptdir}/manager.yaml
kubectl apply -f ${scriptdir}/empty-operator.yaml
kubectl apply -f ${scriptdir}/empty-config.yaml
