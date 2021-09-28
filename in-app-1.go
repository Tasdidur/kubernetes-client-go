package main

import (
	"context"
	"fmt"
	"flag"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)
func main(){
	kubeconfig := flag.String("kubeconfig","/home/office/.kube/config","location to your file")
	config, err := clientcmd.BuildConfigFromFlags("",*kubeconfig)
	if err !=nil {
		// handle errors
		fmt.Printf("error %s building config from flag",err.Error())
		config, err = rest.InClusterConfig()
		if err !=nil {
			fmt.Printf("error %s getting inclusterconfig",err.Error())
		}
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err !=nil {
		// handle error
		fmt.Printf("error %s creating clientset",err.Error())
	}

	ctx := context.Background()

	pods, err :=clientset.CoreV1().Pods("default").List(ctx,metav1.ListOptions{})

	if err != nil {
		//handle error
		fmt.Printf("error %s while listing pods",err.Error())
	}

	for _ , pod := range pods.Items{
		fmt.Printf("%s\n", pod.Name)
	}

}