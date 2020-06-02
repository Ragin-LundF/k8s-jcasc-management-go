# K8S-Management (Go version)

This is a Go implementation of [k8s-jcasc-management](https://github.com/Ragin-LundF/k8s-jcasc-management).

Currently, it is not fully working!

To test it, you have to clone [k8s-jcasc-management](https://github.com/Ragin-LundF/k8s-jcasc-management) and to set the path to this directory as first argument:

```bash
go run k8s-jcasc-mgmt.go -basepath="/path/to/k8s-jcasc-management"
```

## Prerequisites
- Go >= 1.14
- Gpg installed

## Status implementation

Finished:
- Create password for Jenkins user
- Encrypt / decrypt secrets file (global)
- Apply secrets to namespace
- Apply secrets to all namespaces

Open:
- Install Jenkins
- Upgrade Jenkins
- Uninstall Jenkins
- Create project
- Create project for deployment
