package api

import (
	"context"
	"fmt"

	v1 "k8s.io/client-go/kubernetes/typed/apps/v1"

	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/util/retry"
)

type Dep struct {
	DeploymentClient v1.DeploymentInterface
}

//func (dep dep) init() {
//	deploymentsClient := GetClientSet().AppsV1().Deployments(apiv1.NamespaceDefault)
//	dep := dep{
//		deploymentClient: deploymentsClient,
//	}
//}
//var deploymentsClient = GetClientSet().AppsV1().Deployments(apiv1.NamespaceDefault)

func (d Dep) CreateDelployment() {

	//dep := dep{
	//	deploymentClient: deploymentsClient,
	//}

	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: "demo-deployment",
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: intPtr(2),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "demo",
				},
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "demo",
					},
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:  "web",
							Image: "nginx",
							Ports: []apiv1.ContainerPort{
								{
									Name:          "http",
									Protocol:      apiv1.ProtocolTCP,
									ContainerPort: 80,
								},
							},
						},
					},
				},
			},
		},
	}
	result, err := d.DeploymentClient.Create(context.TODO(), deployment, metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Println(result.GetObjectMeta().GetName())
}
func (d Dep) UpdateDeployment() {
	//deploymentsClient := GetClientSet().AppsV1().Deployments(apiv1.NamespaceDefault)
	retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		result, getErr := d.DeploymentClient.Get(context.TODO(), "demo-deployment", metav1.GetOptions{})
		if getErr != nil {
			panic(fmt.Errorf("failed to get latest version of deployment %v ", getErr))
		}
		result.Spec.Replicas = intPtr(3)
		result.Spec.Template.Spec.Containers[0].Image = "nginx:1.13"
		_, updateErr := d.DeploymentClient.Update(context.TODO(), result, metav1.UpdateOptions{})
		return updateErr
	})
	if retryErr != nil {
		panic(fmt.Errorf("update failed %v", retryErr))
	}

}
func (d Dep) GetDeployment() {
	//deploymentClient := GetClientSet().AppsV1().Deployments(apiv1.NamespaceDefault)
	list, err := d.DeploymentClient.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	for _, d := range list.Items {
		fmt.Printf("* %s (%d replicas)\n", d.Name, *d.Spec.Replicas)
	}
}
func (d Dep) DeleteDeployment() {
	//deploymentClient := GetClientSet().AppsV1().Deployments(apiv1.NamespaceDefault)
	deletePolicy := metav1.DeletePropagationForeground
	if err := d.DeploymentClient.Delete(context.TODO(), "demo-deployment", metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	}); err != nil {
		panic(err)
	}
}
func intPtr(n int32) *int32 {
	return &n
}
