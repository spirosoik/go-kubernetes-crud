package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/labstack/gommon/log"
	kubecrud "github.com/spirosoik/go-kubernetes-crud"
)

func main() {
	var kubecontext string
	flag.StringVar(&kubecontext, "kubecontext", "", "The kubecontext to run")
	flag.Parse()

	if kubecontext == "" {
		log.Fatalf("Provide a kubecontext please")
	}
	svc, err := kubecrud.NewService(&kubecontext)

	if err != nil {
		log.Fatal(err)
	}

	name := "test"
	exist, _ := svc.Exist(context.TODO(), name)
	if !exist {
		err = svc.Create(context.TODO(), name)
		if err != nil {
			log.Fatal(err)
		}
	}
	ns, err := svc.Get(context.TODO(), name)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Namespace: %s\n", ns.Name)

	err = svc.Update(context.TODO(), name, map[string]string{
		"testlbl":  "test",
		"testlbl1": "test1",
		"testlbl2": "test2",
	})
	if err != nil {
		log.Fatal(err)
	}
	ns, err = svc.Get(context.TODO(), name)
	fmt.Printf("Labels: %s\n", ns.Labels)

	if err != nil {
		log.Fatal(err)
	}

	err = svc.Delete(context.TODO(), name)
	if err != nil {
		log.Fatal(err)
	}

	exist, _ = svc.Exist(context.TODO(), name)
	if !exist {
		fmt.Printf("Namespace: %s, deleted", name)
	}
}
