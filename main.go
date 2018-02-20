package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	azureRedisClientSet "github.com/amitkr0201/kube-custom-controller/client/clientset/versioned"
	controller "github.com/amitkr0201/kube-custom-controller/controller"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	kubeconfig = flag.String("kubeconfig", "", "Path to a kubeconfig. Only required if out-of-cluster.")
	master     = flag.String("master", "", "The address of the Kubernetes API server. Overrides any value in kubeconfig. Only required if out-of-cluster.")
)

func main() {
	flag.Parse()

	if *kubeconfig == "" {
		log.Printf("No kubeconfig provided. Doing nothing.")
		os.Exit(0)
	}

	cfg, err := clientcmd.BuildConfigFromFlags(*master, *kubeconfig)
	if err != nil {
		log.Printf("Error building kubeconfig: %v", err)
	}

	azureRedisClient, err := azureRedisClientSet.NewForConfig(cfg)
	if err != nil {
		log.Printf("Error building example clientset: %v", err)
	}
	kubeclient, _ := kubernetes.NewForConfig(cfg)
	c := controller.NewController(azureRedisClient, kubeclient)
	stopCh := make(chan struct{})
	defer close(stopCh)
	go c.Run(stopCh)

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGTERM)
	signal.Notify(sigterm, syscall.SIGINT)
	<-sigterm
}
