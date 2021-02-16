# k8s-crd-boilerplate

> :warning: Kubernetes has evolved a lot since I wrote this boilerplate, these days I recommend you look in the directory of: [https://github.com/operator-framework/operator-sdk]()

As I started writing Kubernetes CDR (Custom Resource Definitions) I found there
was a lot of documentation on the details but very little on what was necessary
or not to create a project.

This boilerplate contains only what is needed to bootstrap a CRD.

1. Rename files/packages name from CrdBoilerplate to whatever you are working on
2. Edit types.go and register.go (in 'pkg/apis/crdboilerplate/v1alpha1/')
3. Edit ./hack/update-codegen.sh to reflect the change
4. Generate the code
```
./hack/update-codegen.sh
```
5. Build the controller
```
go build -o list-things cmd/ls-things/main.go
```
6. Edit the CRD manifest (in './artifacts/')
7. Apply it to the cluster and run your binary
```
kubectl apply -f artifacts/cdr.yml
```
8. Add a new test resource
```
kubectl apply -f artifacts/example-thing.yml
```
9. List all the things
```
./ls-things -kubeconfig ~/.kube/config
Found a thing with the name 'a-thing' and the content 'something'
```

I choose not to include the controller daemon code as you can find a very good
and up-to-date example on [github.com/kubernetes/sample-controller](https://github.com/kubernetes/sample-controller/blob/master/controller.go).
