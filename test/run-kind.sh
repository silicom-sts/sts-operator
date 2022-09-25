#!/bin/bash

scriptdir="$( dirname -- "${BASH_SOURCE[0]}"; )";

configDir=${scriptdir}/../config

kubectl create ns openshift-operators

kubectl kustomize ${configDir}/default | kubectl delete -f -

kubectl label nodes kind-control-plane node-role.kubernetes.io/master="" --overwrite
kubectl label nodes kind-control-plane feature.node.kubernetes.io/custom-silicom.sts.devices=true --overwrite
kubectl label nodes kind-control-plane sts.silicom.com/config=bc-1 --overwrite
kubectl label nodes kind-control-plane sts.silicom.com/mode=gm-1 --overwrite
kubectl apply -f https://raw.githubusercontent.com/openshift/cluster-nfd-operator/release-4.10/manifests/4.10/manifests/nfd.openshift.io_nodefeaturediscoveries.yaml

kubectl kustomize ${configDir}/default | kubectl apply -f -
kubectl apply -f ${configDir}/samples/sts_v1alpha1_stsoperatorconfig.yaml
kubectl apply -f ${configDir}/samples/sts_v1alpha1_bc_stsconfig.yaml
kubectl apply -f ${configDir}/samples/sts_v1alpha1_gm_stsconfig.yaml
