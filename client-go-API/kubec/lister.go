package kubec

import (
	"context"
	"fmt"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// func Lister() {
// 	fmt.Println("Running Lister Example")
// 	cs := getKubeHandle()

// 	chck, lk := fields.Everything().RequiresExactMatch("test-pod")
// 	// hk := fields.Everything()
// 	fmt.Println(chck, lk)

// 	// list the pods and create the watch from the server to the cache
// 	listWatch := cache.NewListWatchFromClient(cs.CoreV1().RESTClient(), "pods", "", fields.Everything())
// 	// Pull out the List that implements listwatch
// 	ro, err := listWatch.List(metav1.ListOptions{})
// 	if err != nil {
// 		fmt.Println("err=", err)
// 	}

// 	pods := ro.(*v1.PodList)

// 	fmt.Println("Pods from Lister")
// 	for j, pod := range pods.Items {
// 		fmt.Printf("%d) %v \n", j, pod.Name)
// 	}

// }

func Lister() {
	fmt.Println("Running Lister Example")
	// clientcmd is a subset of client-go pkg
	// client config loading rules to lookup user's kube config
	rules := clientcmd.NewDefaultClientConfigLoadingRules()
	// grab user's kube config
	kubeconfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(rules, &clientcmd.ConfigOverrides{})
	config, err := kubeconfig.ClientConfig()
	if err != nil {
		panic(err)
	}
	clientset := kubernetes.NewForConfigOrDie(config)

	// list the pods using clientset object that calls Pod() method
	listPod := clientset.CoreV1().Pods("")
	// Pull out the List that implements listwatch
	pods, err := listPod.Watch(
		context.Background(),
		metav1.ListOptions{
			// LabelSelector: "app.kubernetes.io/name=nats",
		},
	)
	if err != nil {
		fmt.Println("err=", err)
	}

	go func() {
		for event := range pods.ResultChan() {
			fmt.Printf(
				"Watch Event: %s %s\n",
				event.Type, event.Object.GetObjectKind().GroupVersionKind().Kind,
			)
		}
	}()

	pods.Stop()
	time.Sleep(1 * time.Second)

	// pods := ro.(*v1.PodList)

	// fmt.Println("Pods from Lister")
	// for j, pod := range pods.Items {
	// 	fmt.Printf("%d) %v \n", j, pod.Name)
	// }

}
