# Kubernetes CRUD
A library for running kubernetes CRUD operations.


## Example

You can find a working example how to use it with your own cluster
[here](./examples).

```bash
go run examples/example.go -kubecontext=<your-kube-context>
```

# Usage

Passing nil in the service will try to build a kubernetes client in cluster.

```go
svc := k8crud.NewService(nil)

```

Passing the context you want  will try to build a kubernetes client for the provided context.

```go
kcontext := "my-example"
svc := k8crud.NewService(&kcontext)
```

