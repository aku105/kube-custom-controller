package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	controller "github.com/amitkr0201/kube-custom-controller/controller"
	azureRedisClientSet "github.com/amitkr0201/kube-custom-controller/pkg/client/clientset/versioned"
	"github.com/golang/glog"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	masterURL  string
	kubeconfig string
)

func main() {
	flag.Parse()

	if kubeconfig == "" {
		glog.Fatalf("No kubeconfig provided. Doing nothing.")
		os.Exit(0)
	}

	cfg, err := clientcmd.BuildConfigFromFlags(masterURL, kubeconfig)
	if err != nil {
		glog.Fatalf("Error building kubeconfig: %s", err.Error())
	}

	azureRedisClient, err := azureRedisClientSet.NewForConfig(cfg)
	if err != nil {
		glog.Fatalf("Error building example clientset: %s", err.Error())
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

func init() {
	flag.StringVar(&kubeconfig, "kubeconfig", "", "Path to a kubeconfig. Only required if out-of-cluster.")
	flag.StringVar(&masterURL, "master", "", "The address of the Kubernetes API server. Overrides any value in kubeconfig. Only required if out-of-cluster.")
}
