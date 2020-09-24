# Description
CLI Tool to list all the configured hostnames per namespace on an kubernetes cluster.

## prerequisites
Install kubectl: https://kubernetes.io/docs/tasks/tools/install-kubectl/


### Install and run
CLI Tool to list all the configured hostnames per namespace on an kubernetes cluster.

``` shell
$ export K8S_KUBECONFIG='~/.kube/config' # give full path if ~ gives an error
$ git clone https://github.com/bartvanbenthem/k8s-mgmt.git
$ sudo cp k8s-mgmt/cli-tools/bin/k8mhostname /usr/bin/

$ k8mhostname
```
