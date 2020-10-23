# Description
Multiple packages and a CLI interface to effective administer kubernetes clusters.

## Prerequisites
Install kubectl: https://kubernetes.io/docs/tasks/tools/install-kubectl/


### Install and run
CLI Tool to list all the configured hostnames per namespace on an kubernetes cluster.

``` shell
$ export K8S_KUBECONFIG='~/.kube/config' # give full path if ~ gives an error
$ git clone https://github.com/bartvanbenthem/k8s-administration.git
$ sudo k8s-administration/bin/k8s-administration /usr/bin/

$ k8s-k8s-administration
```

## Functions (in development)
* export all ingress hostnames and corresponding namespaces on the cluster 
* export all assigned admin roles on the cluster 
* export all privelidged containers running on the cluster
... 

