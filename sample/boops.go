// Copyright © 2021 Kris Nóva <kris@nivenly.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
//   ███╗   ██╗ █████╗ ███╗   ███╗██╗
//   ████╗  ██║██╔══██╗████╗ ████║██║
//   ██╔██╗ ██║███████║██╔████╔██║██║
//   ██║╚██╗██║██╔══██║██║╚██╔╝██║██║
//   ██║ ╚████║██║  ██║██║ ╚═╝ ██║███████╗
//   ╚═╝  ╚═══╝╚═╝  ╚═╝╚═╝     ╚═╝╚══════╝
//

package main

import (
	"context"
	"fmt"
	"os"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/kris-nova/naml"
	"k8s.io/client-go/kubernetes"
)

var Version string = "0.0.1"

func main() {
	// Load the application into the NAML registery
	// Note: naml.Register() can be used multiple times.
	//
	naml.Register(NewApp("App", "very serious grown up business application does important beep boops"))

	// Run the generic naml command line program with
	// the application loaded.
	err := naml.RunCommandLine()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

type App struct {
	metav1.ObjectMeta
	description string
	// --------------------
	// Add your fields here
	// --------------------
}

// NewApp will create a new instance of App.
//
// See https://github.com/naml-examples for more examples.
//
// Example: func NewApp(name string, example string, something int) *App
func NewApp(name, description string) *App {
	return &App{
		description: description,
		ObjectMeta: metav1.ObjectMeta{
			Name:            name,
			ResourceVersion: Version,
		},
		// --------------------
		// Add your fields here
		// --------------------
	}
}

func (a *App) Install(client *kubernetes.Clientset) error {
	var err error

	boops648cdbd66c7t4d9Pod := &corev1.Pod{TypeMeta: metav1.TypeMeta{Kind: "Pod",
		APIVersion: "corev1"},
		ObjectMeta: metav1.ObjectMeta{Name: "boops-648cdbd66c-7t4d9",
			GenerateName:               "boops-648cdbd66c-",
			Namespace:                  "default",
			SelfLink:                   "",
			UID:                        "f55ff331-9863-4c78-a3ed-5c599b108894",
			ResourceVersion:            "97197",
			Generation:                 0,
			DeletionGracePeriodSeconds: (*int64)(nil),
			Labels: map[string]string{"app": "boops",
				"pod-template-hash": "648cdbd66c"},
			Annotations: map[string]string(nil),
			OwnerReferences: []metav1.OwnerReference{metav1.OwnerReference{APIVersion: "apps/corev1",
				Kind: "ReplicaSet",
				Name: "boops-648cdbd66c",
				UID:  "d3cb149a-643a-4c9a-9393-bb429040d212",
			}},
			Finalizers:    []string(nil),
			ClusterName:   "",
			ManagedFields: []metav1.ManagedFieldsEntry(nil)},
		Spec: corev1.PodSpec{
			Volumes: []corev1.Volume{corev1.Volume{Name: "kube-api-access-6dv7q",
				VolumeSource: corev1.VolumeSource{HostPath: (*corev1.HostPathVolumeSource)(nil)},
			},
			},
			InitContainers: []corev1.Container(nil),
			Containers: []corev1.Container{corev1.Container{Name: "nginx",
				Image:      "nginx",
				Command:    []string(nil),
				Args:       []string(nil),
				WorkingDir: "",
				Ports:      []corev1.ContainerPort(nil),
				EnvFrom:    []corev1.EnvFromSource(nil),
				Env:        []corev1.EnvVar(nil),
				Resources: corev1.ResourceRequirements{Limits: corev1.ResourceList(nil),
					Requests: corev1.ResourceList(nil)},
				VolumeMounts: []corev1.VolumeMount{corev1.VolumeMount{Name: "kube-api-access-6dv7q",
					ReadOnly:         true,
					MountPath:        "/var/run/secrets/kubernetes.io/serviceaccount",
					SubPath:          "",
					MountPropagation: (*corev1.MountPropagationMode)(nil),
					SubPathExpr:      ""}},
				VolumeDevices:            []corev1.VolumeDevice(nil),
				LivenessProbe:            (*corev1.Probe)(nil),
				ReadinessProbe:           (*corev1.Probe)(nil),
				StartupProbe:             (*corev1.Probe)(nil),
				Lifecycle:                (*corev1.Lifecycle)(nil),
				TerminationMessagePath:   "/dev/termination-log",
				TerminationMessagePolicy: "File",
				ImagePullPolicy:          "Always",
				SecurityContext:          (*corev1.SecurityContext)(nil),
				Stdin:                    false,
				StdinOnce:                false,
				TTY:                      false}},
			EphemeralContainers:          []corev1.EphemeralContainer(nil),
			RestartPolicy:                "Always",
			ActiveDeadlineSeconds:        (*int64)(nil),
			DNSPolicy:                    "ClusterFirst",
			NodeSelector:                 map[string]string(nil),
			ServiceAccountName:           "default",
			AutomountServiceAccountToken: (*bool)(nil),
			NodeName:                     "kind-control-plane",
			HostNetwork:                  false,
			HostPID:                      false,
			HostIPC:                      false,
			ShareProcessNamespace:        (*bool)(nil),
			ImagePullSecrets:             []corev1.LocalObjectReference(nil),
			Hostname:                     "",
			Subdomain:                    "",
			Affinity:                     (*corev1.Affinity)(nil),
			SchedulerName:                "default-scheduler",
			Tolerations: []corev1.Toleration{corev1.Toleration{Key: "node.kubernetes.io/not-ready",
				Operator: "Exists",
				Value:    "",
				Effect:   "NoExecute",
			},
				corev1.Toleration{Key: "node.kubernetes.io/unreachable",
					Operator: "Exists",
					Value:    "",
					Effect:   "NoExecute"}},
			HostAliases:               []corev1.HostAlias(nil),
			PriorityClassName:         "",
			DNSConfig:                 (*corev1.PodDNSConfig)(nil),
			ReadinessGates:            []corev1.PodReadinessGate(nil),
			RuntimeClassName:          (*string)(nil),
			Overhead:                  corev1.ResourceList(nil),
			TopologySpreadConstraints: []corev1.TopologySpreadConstraint(nil),
			SetHostnameAsFQDN:         (*bool)(nil)},
	}

	_, err = client.CoreV1().Pods("default").Create(context.TODO(), boops648cdbd66c7t4d9Pod, metav1.CreateOptions{})
	if err != nil {
		return err
	}

	return err
}

func (a *App) Uninstall(client *kubernetes.Clientset) error {
	var err error

	err = client.CoreV1().Pods("default").Delete(context.TODO(), "boops648cdbd66c7t4d9", metav1.DeleteOptions{})
	if err != nil {
		return err
	}

	return err
}

func (a *App) Description() string {
	return a.description
}

func (a *App) Meta() *metav1.ObjectMeta {
	return &a.ObjectMeta
}