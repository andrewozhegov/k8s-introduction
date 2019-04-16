# Introducion into CI/CD & Kubernetes

Step-by-step process of deployment a simple Go web-service into Kubernetes

![docker and kubernetes](https://dotmesh.io/assets/img/docker_kubernetes.png)


## Step 1. Pre-commit-hooks

A set of configurable code quality checks to perform before committing.
Specify hooks in `.pre-commit-config.yaml` file:

```
-   repo: git://github.com/dnephin/pre-commit-golang      # repo with hooks
    sha: 471a7c123ea7a3b776ff018edf00066947873a94         # revision
    hooks:
    -   id: go-fmt                                        # single codestyle analyzer
    -   id: go-vet                                        # check for any errors in packages
    -   id: go-lint                                       # check for any errors in code
```

## Step 2. Simple Go web-service architecture

### Router
### Logging
### Healthchecks: liveness & readiess
### Graceful shutdown

## Step 3. Tests

## Step 4. Managing dependencies with Glide

## Step 5. CI/CD process in Makefile

### clean
### vendor
### build
### container
### push
### run
### deploy

## Step 6. Configuring & Versioning

## Step 7. Deploy into minikube

