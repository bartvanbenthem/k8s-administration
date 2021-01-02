# Description
Print all ingress hostnames and corresponding namespaces and contexts.

## Prerequisites
Install kubectl: https://kubernetes.io/docs/tasks/tools/install-kubectl/


### Install and run
CLI Tool to list all the configured hostnames per namespace on an kubernetes cluster.

``` shell
$ export K8S_KUBECONFIG='~/.kube/config' # give full path if ~ gives an error
$ git clone https://github.com/bartvanbenthem/k8s-listners.git
$ sudo k8s-listners/bin/k8s-listners /usr/bin/

$ k8s-listners
```
