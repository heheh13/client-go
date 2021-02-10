package api

type Resource struct {
	filePath string
}

func (r Resource) Create() {
	r.filePath = "i want here? "
}

//deployment:?
// create from files
// kc create deployment <deploymentName> <deployment_image
// kc describe
// kc edit
