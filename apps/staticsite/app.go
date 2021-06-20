//
// Copyright © 2021 Kris Nóva <kris@nivenly.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, softwar
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
//    ███╗   ██╗ ██████╗ ██╗   ██╗ █████╗
//    ████╗  ██║██╔═████╗██║   ██║██╔══██╗
//    ██╔██╗ ██║██║██╔██║██║   ██║███████║
//    ██║╚██╗██║████╔╝██║╚██╗ ██╔╝██╔══██║
//    ██║ ╚████║╚██████╔╝ ╚████╔╝ ██║  ██║
//    ╚═╝  ╚═══╝ ╚═════╝   ╚═══╝  ╚═╝  ╚═╝

package staticsite

import (
	"context"
	"fmt"
	yamyams "github.com/kris-nova/yamyams/pkg"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"

	appsv1 "k8s.io/api/apps/v1"

	v1 "k8s.io/api/apps/v1"

	apiv1 "k8s.io/api/core/v1"
)

const (
	name string = "yamyams-static-site"
)

// StaticSite represents a static site.
type StaticSite struct {
	resources []interface{}
	meta      *yamyams.DeployableMeta
	namespace string
}

func New(namespace string) *StaticSite {
	return &StaticSite{
		namespace: namespace,
		meta: &yamyams.DeployableMeta{
			Name:        "Example Static Website",
			Version:     "0.0.1",
			Command:     "staticsite",
			Description: "A simple static website application",
		},
	}
}

func (v *StaticSite) Install(client *kubernetes.Clientset) error {

	containerImage := "krisnova/yamyams-static-site"
	//containerPort := 80
	labels := map[string]string{
		"yamyams": "yamyams",             // Add this to every app so you can kubectl get po -l yamyams=yamyams
		"app":     "yamyams-static-site", // Add this to this application so you can kubectl get po -l name=yamyams-static-site
	}

	deployment := &v1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: yamyams.I32p(2),
			Selector: &metav1.LabelSelector{
				MatchLabels: labels,
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: labels,
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:  name,
							Image: containerImage,
							Ports: []apiv1.ContainerPort{
								{
									Name:          "http",
									Protocol:      apiv1.ProtocolTCP,
									ContainerPort: 80, // It's fine to hard code this because you would have done that shit anyway in YAML
								},
							},
						},
					},
				},
			},
		},
	}

	updatedDeployment, err := client.AppsV1().Deployments(v.namespace).Create(context.TODO(), deployment, metav1.CreateOptions{})
	if err != nil {
		v.resources = append(v.resources, deployment)
		return fmt.Errorf("unable to install in Kubernetes: %v", err)
	}
	v.resources = append(v.resources, updatedDeployment)
	return nil
}

func (v *StaticSite) Uninstall(client *kubernetes.Clientset) error {
	return client.AppsV1().Deployments(v.namespace).Delete(context.TODO(), name, metav1.DeleteOptions{})
}

func (v *StaticSite) Resources() []interface{} {
	return v.resources
}

func (v *StaticSite) About() *yamyams.DeployableMeta {
	return v.meta
}
