package main

import (
	"context"
	"flag"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)
func main(){
	kubeconfig := flag.String("kubeconfig","/home/office/.kube/config","location to your file")
	config, err := clientcmd.BuildConfigFromFlags("",*kubeconfig)
	if err !=nil {
		// handle errors
		fmt.Printf("error %s building config from flag",err.Error())
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err !=nil {
		// handle error
		fmt.Printf("error %s creating clientset",err.Error())
	}

	ctx := context.Background()

	api := clientset.CoreV1()

	pods, err := api.Pods("default").List(ctx,metav1.ListOptions{})

	if err != nil {
		//handle error
		fmt.Printf("error %s while listing pods",err.Error())
	}

	for _ , pod := range pods.Items{
		fmt.Printf("%s\n", pod.Name)
	}

}