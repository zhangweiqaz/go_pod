# Overview

just to execute the command: 
kubectl exec -it podName [-c ContainerName] [-n NameSpace] -- su - userName

# Install

Download the binary file kubectl-go, and put into /usr/local/bin. It means that command must exist in your PATH.

# Usage
kubectl go podName [-c ContainerName] [-n NameSpace] [-u UserName]
```bash
kubectl go -h
kubectl exec in pod with username. For example:
kubectl go pod_name

Usage:
  go [flags]

Flags:
  -c, --containerName string   containerName
  -h, --help                   help for go
  -n, --namespace string       namespace
  -u, --username string        username, this user must exist in image, default: dev
```

# Attention

* the user must exist in the image, you can add "RUN useradd dev" to create one
* su must exist in the image
