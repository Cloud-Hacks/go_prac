package main

import (
	"context"
	"fmt"

	// corev1 "k8s.io/api/core/v1"
	c1 "github.com/afzalbin64/client-go/kubec"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
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

	nodeList, err := clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	for _, n := range nodeList.Items {
		fmt.Println(n.Labels)
	}

	defer context.Canceled.Error()

	// URL := "http://127.0.0.1:8080" + "/api/v1/nodes?watch=true"

	// resp, _ := http.Get(URL)

	// fmt.Println(resp)

	// newPod := &corev1.Pod{
	// 	ObjectMeta: metav1.ObjectMeta{
	// 		Name: "test-pod",
	// 	},
	// 	Spec: corev1.PodSpec{
	// 		Containers: []corev1.Container{
	// 			{Name: "resmgmnt-api", Image: "afzal442/resmgmnt-api:v1"},
	// 		},
	// 	},
	// }

	// pod, err := clientset.CoreV1().Pods("default").Create(context.Background(), newPod, metav1.CreateOptions{})
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(pod)

	// c1.Informer()
	// c1.CrudOperation()
	c1.Lister()
	// c1.WatchTypedClient()
	// c1.Workqueue_example()
}
