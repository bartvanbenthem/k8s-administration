#!/bin/bash

# initialize go modules
go mod init github.com/bartvanbenthem/k8s-hostname
# get the correct go-client module
go get k8s.io/client-go@kubernetes-1.16.7