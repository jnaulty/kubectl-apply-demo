# Kubectl Apply Demo

Here's a small demo app and static files to illustrate why you should not run `kubectl apply` commands for random `https` endpoints.

## Kubectl Apply Info

`kubectl apply` is used to `Apply a configuration change to a resource from a file or stdin.`

### Problem

`kubectl` documentation does not warn about the possible dangers of applying random yaml files from the internet.

An example from [k8s docs on "Managing Resources"](https://kubernetes.io/docs/concepts/cluster-administration/manage-deployment/) might have us believe that applying random yaml files from the internet is safe:
`kubectl apply -f https://k8s.io/examples/application/nginx-app.yaml`


### Solution

Applying random things from the internet into a cluster seems like a good way for bad things to happen. Instead, I suggest the following pattern of "get, inspect, apply"

**Get**: `wget https://k8s.io/examples/application/nginx-app.yaml`

**Inspect**: `nginx-app.yaml`

**Apply**: `kubectl apply -f nginx-app.yaml`



