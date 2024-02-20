
package main

import (
    "context"
    "fmt"
    "k8s.io/client-go/kubernetes"
    "k8s.io/client-go/tools/clientcmd"
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func main() {
    config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
    if err != nil {
        panic(err.Error())
    }

    clientset, err := kubernetes.NewForConfig(config)
    if err != nil {
        panic(err.Error())
    }

    pods, err := clientset.CoreV1().Pods("default").List(context.TODO(), metav1.ListOptions{})
    if err != nil {
        panic(err.Error())
    }

    fmt.Println("Listing pods in the default namespace:")
    for _, pod := range pods.Items {
        fmt.Printf("%s\t%s\t%s\n", pod.Name, pod.Status.Phase, pod.Spec.Containers[0].Image)
    }
}
