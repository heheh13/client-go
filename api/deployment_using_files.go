package api

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	v1 "k8s.io/client-go/kubernetes/typed/apps/v1"

	"github.com/kr/pretty"

	appsv1 "k8s.io/api/apps/v1"
)

type Resource struct {
	//holds on the file path of the wanted api object
	FilePath string
	// type of the api obeject like depl,statefulset,etc
	Kind appsv1.Deployment
	// such as depl client, statefulset client
	Clientset v1.DeploymentInterface
}

func (r Resource) Create() {
	file, err := os.Open(r.FilePath)
	fmt.Println(r.FilePath)
	if err != nil {
		panic(err.Error())
	}
	dec := json.NewDecoder(file)
	dec.Decode(&r.Kind)
	pretty.Println(r.Kind)
	result, err := r.Clientset.Create(context.TODO(), &r.Kind, metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Println(result.GetObjectMeta().GetName())
}
