#!/bin/sh

# exports the configuration of the k8s into a DIR.
# this can be helpful if you need to "debug" the configuration.

# https://artem.services/?p=1389&lang=en

DIR='k8s-manifests/namespaces'

mkdir -p $DIR
kubectl version -o=yaml         > $DIR/../k8s-version.yaml
kubectl get nodes -o=yaml       > $DIR/../k8s-nodes.yaml
kubectl get nodes -o=wide       > $DIR/../k8s-nodes.wide
kubectl describe nodes          > $DIR/../k8s-nodes.describe

for NAMESPACE in $(kubectl get -o=name namespaces | cut -d '/' -f2)
do
    for TYPE in $(kubectl get -n $NAMESPACE -o=name pvc,configmap,serviceaccount,secret,ingress,service,deployment,statefulset,hpa,job,cronjob)
    do
        mkdir -p $(dirname $DIR/$NAMESPACE/$TYPE)
        echo " ==> kubectl get -n $NAMESPACE -o=yaml $TYPE"
        kubectl get -n $NAMESPACE -o=yaml $TYPE > $DIR/$NAMESPACE/$TYPE.yaml
    done
done
