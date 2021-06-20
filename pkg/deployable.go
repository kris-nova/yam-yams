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
//

package yamyams

import "k8s.io/client-go/kubernetes"

// Deployable is an interface that can be implemented
// for deployable applications.
type Deployable interface {

	// Install will attempt to install in Kubernetes
	Install(client *kubernetes.Clientset) error

	// Uninstall will attempt to uninstall in Kubernetes
	Uninstall(client *kubernetes.Clientset) error

	// Resources returns untyped struct{}s which represent your application.
	// This is the first concrete point in which we realize that what we are doing, has no business being "generic".
	Resources() []interface{}

	//YAML() []byte
	//JSON() []byte
}
