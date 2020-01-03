# Kubectl Apply Demo

Here's a small demo app and static files to illustrate why you should not run `kubectl apply` commands for random `https` endpoints.

## Demo

1. Clone this repo: `git clone https://github.com/jnaulty/kubectl-apply-demo.git`
2. `cd` into repo `cd kubectl-apply-demo`
3. Build binary `go build` or download from release page `wget https://github.com/jnaulty/kubectl-apply-demo/releases/download/v0.1.0/kubectl-apply-demo`
(`cf7dcfc0c9c37d3986361acdee1d4e5172350a85e39e6b9db9f76600cc7218d0` is the sha256sum of this binary)
4. Run server in separate terminal or background `./kubectl-apply-demo`
(You might need to make it executable if you downloaded it from github in Step 3, `chmod u+x kubectl-apply-demo`)
5. Run `curl localhost:8080/harmless.yaml`
6. Run `kubectl apply -f localhost:8080/harmless.yaml --dry-run -o yaml`
7. Spot the difference

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

At the very least, do something like `kubectl apply -f <my-possibly-dangerous-url> --dry-run -o yaml` and inspect the output in the console.



