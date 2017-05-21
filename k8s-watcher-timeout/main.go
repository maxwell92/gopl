package main

import (
	yceutils "app/backend/controller/yce/utils"
	"k8s.io/kubernetes/pkg/api"
	"log"
	"reflect"
)

func main() {

	cli, _ := yceutils.CreateK8sClient("172.21.1.11:8080")
	watcher1, _ := cli.Pods(api.NamespaceAll).Watch(api.ListOptions{})
	watcher2, _ := cli.Pods(api.NamespaceAll).Watch(api.ListOptions{Watch: true})
	watcher3, _ := cli.Pods(api.NamespaceAll).Watch(api.ListOptions{TimeoutSeconds: yceutils.Int64Ptr(180)})
	watcher4, _ := cli.Pods(api.NamespaceAll).Watch(api.ListOptions{Watch: true, TimeoutSeconds: yceutils.Int64Ptr(180)})

	eventCh1 := watcher1.ResultChan()
	eventCh2 := watcher2.ResultChan()
	eventCh3 := watcher3.ResultChan()
	eventCh4 := watcher4.ResultChan()

	for {
		select {
		case e1 := <-eventCh1:
			{
				log.Printf("e1 %s:%s\n", e1.Type, reflect.TypeOf(e1.Object).String())
			}
		case e2 := <-eventCh2:
			{
				log.Printf("e2 %s:%s\n", e2.Type, reflect.TypeOf(e2.Object).String())
			}
		case e3 := <-eventCh3:
			{
				log.Printf("e3 %s:%s\n", e3.Type, reflect.TypeOf(e3.Object).String())
			}
		case e4 := <-eventCh4:
			{
				log.Printf("e4 %s:%s\n", e4.Type, reflect.TypeOf(e4.Object).String())
			}
		}
	}
}
