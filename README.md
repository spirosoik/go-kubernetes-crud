# Kubernetes CRUD
A library for running kubernetes CRUD operations.


## Example

You can find a working example how to use it with your own cluster
[here](./examples).

```bash
go run examples/example.go -kubecontext=<your-kube-context>
```

## Usage

Passing nil in the service will try to build a kubernetes client in cluster.

```go
svc := kubecrud.NewService(nil)

```

Passing the context you want  will try to build a kubernetes client for the provided context.

```go
kcontext := "my-example"
svc := kubecrud.NewService(&kcontext)
```

## Development

The library is currently compatible with golang version from 1.13+.

Use a [minikube](https://kubernetes.io/docs/tasks/tools/install-minikube/) or [microk8](https://microk8s.io/) for local development with a Kubernetes cluster.

The library uses Go modules so it can be located outside the Go path.

```bash
# Setup with extra tools
make setup

# Run linter for Go
make lint

# Run the example
make example
```