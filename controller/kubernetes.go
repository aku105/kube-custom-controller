package controller

import (
	"fmt"
	"log"
	"time"

	"github.com/amitkr0201/kube-custom-controller/pkg/apis/azurerediscontroller/v1alpha1"
	azureredisclientset "github.com/amitkr0201/kube-custom-controller/pkg/client/clientset/versioned"
	azureredisinformer "github.com/amitkr0201/kube-custom-controller/pkg/client/informers/externalversions"
	azureredisconfiginformer "github.com/amitkr0201/kube-custom-controller/pkg/client/informers/externalversions/azurerediscontroller/v1alpha1"

	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
)

// Execute Executes the controller
func Execute() {
	fmt.Println("I am printed.")
}

// Controller gives controller
type Controller struct {
	crdclientset azureredisclientset.Interface
	queue        workqueue.RateLimitingInterface
	informer     azureredisconfiginformer.AzureRedisInformer
	kubeclient   kubernetes.Interface
}

// NewController Create new controller
func NewController(crdclientset azureredisclientset.Interface, kubeclient kubernetes.Interface) *Controller {
	fmt.Println("new Controller called.")

	azureRedisInfromersfactory := azureredisinformer.NewSharedInformerFactory(crdclientset, time.Second*30)
	azureRedises := azureRedisInfromersfactory.Azurerediscontroller().V1alpha1().AzureRedises()
	configsqueue := workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter())
	log.Println("Setting up event handlers")

	azureRedises.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			key, err := cache.MetaNamespaceKeyFunc(obj)
			if err == nil {
				configsqueue.Add(key)
			}
		},

		UpdateFunc: func(old, new interface{}) {
			newDepl := new.(*v1alpha1.AzureRedis)
			oldDepl := old.(*v1alpha1.AzureRedis)
			if newDepl.ResourceVersion == oldDepl.ResourceVersion {
				// Periodic resync will send update events for all known Resources.
				// Two different versions of the same CRDS (configgit kind) will always have different RVs.
				return
			}

			key, err := cache.MetaNamespaceKeyFunc(new)
			if err == nil {
				configsqueue.Add(key)
			}
		},
	})

	controller := &Controller{
		crdclientset: crdclientset,
		informer:     azureRedises,
		queue:        configsqueue,
		kubeclient:   kubeclient,
	}

	return controller
}

// Run runs the controller
func (c *Controller) Run(stopCh <-chan struct{}) {
	defer utilruntime.HandleCrash()
	defer c.queue.ShutDown()

	log.Println("Starting AzureRedis controller")

	go c.informer.Informer().Run(stopCh)

	if !cache.WaitForCacheSync(stopCh, c.HasSynced) {
		utilruntime.HandleError(fmt.Errorf("Timed out waiting for caches to sync"))
		return
	}

	log.Println("AzureRedis controller synced and ready")

	wait.Until(c.runWorker, time.Second, stopCh)
}

// HasSynced has it synced yet?
func (c *Controller) HasSynced() bool {
	return c.informer.Informer().HasSynced()
}

func (c *Controller) runWorker() {
	for c.processNextItem() {
	}
}

func (c *Controller) processNextItem() bool {
	key, quit := c.queue.Get()
	if quit {
		return false
	}
	defer c.queue.Done(key)

	err := c.processItem(key.(string))
	if err == nil {
		// No error, reset the ratelimit counters
		c.queue.Forget(key)
	} else if c.queue.NumRequeues(key) < 5 {
		log.Printf("Error processing %s (will retry): %v", key, err)
		c.queue.AddRateLimited(key)
	} else {
		// err != nil and too many retries
		log.Printf("Error processing %s (giving up): %v", key, err)
		c.queue.Forget(key)
		utilruntime.HandleError(err)
	}

	return true
}

func (c *Controller) processItem(key string) error {
	// c.logger.Infof("Processing change to Pod %s", key)

	Obj, exists, err := c.informer.Informer().GetIndexer().GetByKey(key)
	if err != nil {
		return fmt.Errorf("Error fetching object with key %s from store: %v", key, err)
	}

	if !exists {

		fmt.Println("CRD deleted - ", key)
		return nil
	}
	log.Println("Creating a new", Obj.(*v1alpha1.AzureRedis).Name, "in resource group", Obj.(*v1alpha1.AzureRedis).Spec.ResourceGroup)
	c.ProcessConfig(*Obj.(*v1alpha1.AzureRedis))
	return nil
}

// ProcessConfig Process config for CRD
func (c *Controller) ProcessConfig(Obj v1alpha1.AzureRedis) {

	log.Println("Processing request for", Obj.Name, "in Resource group:", Obj.Spec.ResourceGroup)
	return

	// directory := "/tmp/gitworkspace/" + Obj.GetObjectMeta().GetResourceVersion() + "/*"
	// log.Println("Directory Path - ", directory)
	// os.RemoveAll(directory)
	// _, err := git.PlainClone(directory, false, &git.CloneOptions{
	// 	URL:               GitUrl,
	// 	RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
	// })

	// if err != nil {
	// 	os.RemoveAll(directory)
	// 	log.Println("Failed to checkout ", err)
	// 	return
	// }
	// log.Println("Removing .git folder ")
	// os.RemoveAll(directory + ".git")
	// Data := make(map[string]string)

	// _ = filepath.Walk(directory, func(path string, f os.FileInfo, err error) error {

	// 	b, err := ioutil.ReadFile(path)
	// 	filecontent := string(b)
	// 	Data[filepath.Base(path)] = filecontent
	// 	return nil
	// })

	// cm := &k8v1.ConfigMap{
	// 	ObjectMeta: metav1.ObjectMeta{
	// 		Name:      Obj.Name,
	// 		Namespace: Obj.Namespace,
	// 	},

	// 	Data: Data,
	// }
	// result, err := c.kubeclient.CoreV1().ConfigMaps("default").Create(cm)
	// if err != nil {
	// 	log.Println("Error creating configmap ", err)
	// 	os.RemoveAll(directory)
	// 	return
	// }

	// log.Println("Created configmap from Git", result.GetObjectMeta().GetName())
	// os.RemoveAll(directory)
}
