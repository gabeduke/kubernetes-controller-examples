package main

import (
    "context"
    "fmt"
    "k8s.io/apimachinery/pkg/apis/meta/v1"
    "k8s.io/apimachinery/pkg/watch"
    "k8s.io/client-go/kubernetes"
    "k8s.io/client-go/tools/clientcmd"
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

    watchInterface, err := clientset.CoreV1().Pods("default").Watch(context.TODO(), v1.ListOptions{})
    if err != nil {
        panic(err.Error())
    }
    fmt.Println("Watching for Pod changes in the default namespace:")
    for event := range watchInterface.ResultChan() {
        pod := event.Object.(*v1.Pod)
        fmt.Printf("Event: %v Pod: %s\n", event.Type, pod.Name)
    }
}
