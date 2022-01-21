package kubec

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func CrudOperation() {
	fmt.Println("Running CRUD Example")
	// get handled to K8s
	cs := getKubeHandle()

	// Interect with K8s client-go V1 API and get list of pods of any namespace
	// In the ListOptions we can define labels such as ns, db etc
	// to filter the list
	pods, err := cs.CoreV1().Pods("").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		fatal(fmt.Sprintf("error getting list of pods: %v", err))
	}

	fmt.Println("## Pods ##")
	for i, pod := range pods.Items {
		fmt.Printf("%d) %v \n", i, pod.Name)
	}
}
