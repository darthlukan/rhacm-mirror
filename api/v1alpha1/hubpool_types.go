/*
Copyright 2022 Brian Tomlinson.

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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CredentialSpec defines the credential structure
type CredentialSpec struct {
	Secret string `json:"secret,omitempty"`
}

// BackupLocationSpec defines the remote backup location
type BackupLocationSpec struct {
	URL         string         `json:"url,omitempty"`
	Protocol    string         `json:"protocol,omitempty"`
	Credentials CredentialSpec `json:"credentials,omitempty"`
}

// HubPoolSpec defines the desired state of HubPool
type HubPoolSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Pool defines the list of hubs in the pool by 'hub_name: api.hub_name.*'
	Pool map[string]string `json:"pool,omitempty"`
	// InitialLeader defines the hub desired as initial lead hub.
	// IMPORTANT! This is purely a RHACM Mirror configuration and DOES NOT configure RHACM itself!
	InitialLeader string `json:"initial_leader,omitempty"`
	// PingPeriod defines the time between pings to hubs in spec.pool
	PingPeriod string `json:"ping_period,omitempty"`
	// PingFailsAcceptable defines the number of failed pings before RHACM Mirror begins failover activities
	PingFailsAcceptable int `json:"ping_fails_acceptable,omitempty"`
	// BackupLocation defines the remote location containing RHACM backups
	BackupLocation BackupLocationSpec `json:"backup_location,omitempty"`
	// FailoverAlertEmail defines the email address to be notified when failover activities have ocurred
	FailoverAlertEmail string `json:"failover_alert_email,omitempty"`
}

// HubPoolStatus defines the observed state of HubPool
type HubPoolStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// HubPool is the Schema for the hubpools API
type HubPool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   HubPoolSpec   `json:"spec,omitempty"`
	Status HubPoolStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// HubPoolList contains a list of HubPool
type HubPoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HubPool `json:"items"`
}

func init() {
	SchemeBuilder.Register(&HubPool{}, &HubPoolList{})
}
