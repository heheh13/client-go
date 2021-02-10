package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/kr/pretty"

	appsv1 "k8s.io/api/apps/v1"
)

var (
	deplJson = "./file.json"
)

func main() {
	//cmd.Execute()
	file, err := os.Open(deplJson)
	if err != nil {
		panic(err)
	}
	dec := json.NewDecoder(file)
	//var dep extensionsv1beta1.Deployment
	var dep appsv1.Deployment
	dec.Decode(&dep)
	fmt.Println(dep)
	//dec.Decode(&dep)
	//pretty.Println(dep)
	pretty.Println(&dep)

}
