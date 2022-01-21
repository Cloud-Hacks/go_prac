package kubec

import (
	"fmt"
	"time"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/client-go/tools/cache"
)

func Informer() {
	fmt.Println("Running Informer Example")
	cs := getKubeHandle()

	// listing the pods for every field and creating a watcher from cache
	listWatch := cache.NewListWatchFromClient(cs.CoreV1().RESTClient(), "pods", "", fields.Everything())

	_, controller := cache.NewInformer(listWatch, &v1.Pod{}, time.Second*5, cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			pod := obj.(*v1.Pod)
			fmt.Println("Pod Added:", pod.Name)
		},
		DeleteFunc: func(obj interface{}) {
			pod := obj.(*v1.Pod)
			fmt.Println("Pod Deleted:", pod.Name)
		},
	})

	stop := make(chan struct{})
	controller.Run(stop)
}
