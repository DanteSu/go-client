package main

import (
	"context"
	"flag"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
)

func main() {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create a RESTclient
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// get all of the pods under kube-system namespace
	podList, err := clientset.CoreV1().Pods("kube-system").List(context.TODO(), metav1.ListOptions{Limit: 100})
	if err != nil {
		panic(err.Error())
	}

	// 输出Pod信息
	fmt.Printf("NAMESPACE\tSTATUS\t\tNAME\n")
	for _, pod := range podList.Items {
		fmt.Printf("%-16s %-16s %s\n", pod.Namespace, pod.Status.Phase, pod.Name)
	}
}
