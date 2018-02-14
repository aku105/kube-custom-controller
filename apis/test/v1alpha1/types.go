package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzureRedis is a test crd.
type AzureRedis struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Specification of the desired behavior of the workspace.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#spec-and-status
	Spec AzureRedisSpec `json:"spec"`

	// Specification of the desired behavior of the workspace.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#spec-and-status
	Status AzureRedisStatus `json:"status"`
}

// AzureRedisSpec is the spec for a AzureRedis resource.
type AzureRedisSpec struct {
	// Azure region where the Redis instance exists.
	// All Azure regions can be found at: https://azure.microsoft.com/en-us/regions/
	Location string `json:"location,omitempty"`

	// SKU type of Azure Redis instance. Allowed values: Basic, Premium, Standard.
	// More info: https://azure.microsoft.com/en-us/pricing/details/cache/
	SKU string `json:"sku,omitempty"`

	// Size of Azure Redis instance.
	// All valid values at: https://azure.microsoft.com/en-us/pricing/details/cache/
	Size string `json:"size,omitempty"`

	// Resource group in which Azure Redis instance is in.
	ResourceGroup string `json:"resourceGroup,omitempty"`

	// Name of the Azure Redis instance. Must be unique accross an Azure region.
	Name string `json:"name"`
}

// AzureRedisStatus is the status for AzureRedis CRD
type AzureRedisStatus struct {
	// Current status of Azure Redis instance.
	// Values: Submitted, In progress, Done, Failed
	CreationStatus string `json:"creationStatus"`
	Message        string `json:"message,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzureRedisList is a list of AzureRedis resources
type AzureRedisList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []AzureRedis `json:"items"`
}
