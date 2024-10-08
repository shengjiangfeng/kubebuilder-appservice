/*
Copyright 2021.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.
// 重新定义service, 改成遍历，加上serviceType

// AppServiceSpec defines the desired state of AppService
type AppServiceSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Replicas            *int32                      `json:"replicas"`
	Image               string                      `json:"image"`
	Resources           corev1.ResourceRequirements `json:"resources,omitempty"`
	GatewayRegisterInfo GatewayRegisterInfo         `json:"gateway_register_info"`
	Envs                []corev1.EnvVar             `json:"envs,omitempty"`
	Ports               []corev1.ServicePort        `json:"ports,omitempty"`
}

// AppServiceStatus defines the observed state of AppService
type AppServiceStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// 这里可以加入自己想跟踪的状态，这个demo就不加了
	// Important: Run "make" to regenerate code after modifying this file

	appsv1.DeploymentStatus `json:",inline"`
}

type GatewayRegisterInfo struct {
	FeatureEnabled bool   `json:"feature_enabled,omitempty"`
	RegisterUrl    string `json:"register_url,omitempty"`
	ServicePath    string `json:"service_path,,omitempty"` // 服务http访问路径,管理网关注册
	Port           int    `json:"port,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// AppService is the Schema for the appservices API
type AppService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AppServiceSpec   `json:"spec,omitempty"`
	Status AppServiceStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// AppServiceList contains a list of AppService
type AppServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AppService `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AppService{}, &AppServiceList{})
}
