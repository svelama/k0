package main

import (
	"context"
	"flag"
	"fmt"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {
	//config, _ := rest.InClusterConfig()

	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}

	clientset, _ := kubernetes.NewForConfig(config)
	events, _ := clientset.CoreV1().Events("default").List(context.TODO(), metav1.ListOptions{TypeMeta: metav1.TypeMeta{Kind: "Node"}})
	for _, item := range events.Items {
		fmt.Println(item.CreationTimestamp, item.Name, item.Namespace, item.Kind, item.ReportingInstance, item.Source, item.Reason, item.Message)
		fmt.Println()
	}
}
