package main

import (
	"context"
	"flag"
	"fmt"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"log"
)

func main(){
	kubeconfig := flag.String("kubeconfig","/home/office/.kube/config","cofig file path")
	//fmt.Println(*config)
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		log.Fatal(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}
	api := clientset.CoreV1()


	var ns, label, field string
	flag.StringVar(&ns, "namespace", "", "namespace")
	flag.StringVar(&label, "l", "", "Label selector")
	flag.StringVar(&field, "f", "", "Field selector")
	flag.Parse()

	listoptions := metav1.ListOptions{LabelSelector: label,FieldSelector: field}

	watcher , err := api.Pods(ns).Watch(context.Background(),listoptions)
	if err !=nil{
		log.Fatal(err)
	}
	ch := watcher.ResultChan()

	template := "%-40s%-15s\n"
	fmt.Printf(template, "NAME", "STATUS")
	for event := range ch {
		pod, ok := event.Object.(*v1.Pod)
		if !ok {
			log.Fatal("unexpected type")
		}
		fmt.Printf(template,pod.Name,string(pod.Status.Phase))
	}

	//-------------------------------------------------------------------------------------
	//pods, err := api.Pods(ns).List(context.Background(),listoptions)
	//if err !=nil{
	//	log.Fatal(err)
	//}
	//
	//template := "%-40s%-15s\n"
	//fmt.Printf(template, "NAME", "STATUS")
	//for _, pod := range pods.Items {
	//	fmt.Printf(
	//		template,
	//		pod.Name,
	//		string(pod.Status.Phase))
	//}
	//-------------------------------------------------------------------------------------
	//for _,pod := range pods.Items{
	//	fmt.Printf("%s\n", pod.Name)
	//}

}