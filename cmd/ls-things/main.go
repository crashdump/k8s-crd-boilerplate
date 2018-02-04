package main

import (
	"flag"
	"fmt"

	"github.com/golang/glog"
	"k8s.io/client-go/tools/clientcmd"

	boilerplateset "github.com/crashdump/k8s-crd-boilerplate/pkg/client/clientset/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	kuberconfig = flag.String("kubeconfig", "", "Path to a kubeconfig. Only required if out-of-cluster.")
	master      = flag.String("master", "", "The address of the Kubernetes API server. Overrides any value in kubeconfig. Only required if out-of-cluster.")
)

func main() {
	flag.Parse()

	cfg, err := clientcmd.BuildConfigFromFlags(*master, *kuberconfig)
	if err != nil {
		glog.Fatalf("Error building kubeconfig: %v", err)
	}

	boilerplateClient, err := boilerplateset.NewForConfig(cfg)
	if err != nil {
		glog.Fatalf("Error building boilerplate clientset: %v", err)
	}

	list, err := boilerplateClient.CrdboilerplateV1alpha1().Things("default").List(metav1.ListOptions{})
	if err != nil {
		glog.Fatalf("Error listing all things: %v", err)
	}

	for _, thing := range list.Items {
		fmt.Printf("Found a thing with the name '%s' and the content '%s'\n", thing.Spec.Name, thing.Spec.Content)
	}
}
