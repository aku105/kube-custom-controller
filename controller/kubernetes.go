package controller

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	examplecomclientset "github.com/amitkr0201/kube-custom-controller/client/clientset/versioned"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// Execute Executes the controller
func Execute() {
	fmt.Println("I am printed.")

	kubeconfig := flag.String("kubeconfig", "", "Path to a kubeconfig. Only required if out-of-cluster.")
	master := flag.String("master", "", "The address of the Kubernetes API server. Overrides any value in kubeconfig. Only required if out-of-cluster.")

	flag.Parse()

	cfg, err := clientcmd.BuildConfigFromFlags(*master, *kubeconfig)

	if err != nil {
		fmt.Printf("Error building kubeconfig: %v", err)
	}

	exampleClient, err := examplecomclientset.NewForConfig(cfg)
	if err != nil {
		fmt.Printf("Error building example clientset: %v", err)
	}
	kubeclient, _ := kubernetes.NewForConfig(cfg)
	c := controller.NewController(exampleClient, kubeclient)
	stopCh := make(chan struct{})
	defer close(stopCh)
	go c.Run(stopCh)

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGTERM)
	signal.Notify(sigterm, syscall.SIGINT)
	<-sigterm
}
