/*
Copyright (c) 2019 Red Hat, Inc.

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

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1 // github.com/openshift-online/uhc-sdk-go/pkg/client/clustersmgmt/v1

import (
	"github.com/openshift-online/uhc-sdk-go/pkg/client/helpers"
)

// clusterRegistrationData is the data structure used internally to marshal and unmarshal
// objects of type 'cluster_registration'.
type clusterRegistrationData struct {
	SubscriptionID *string "json:\"subscription_id,omitempty\""
	ExternalID     *string "json:\"external_id,omitempty\""
}

// MarshalClusterRegistration writes a value of the 'cluster_registration' to the given target,
// which can be a writer or a JSON encoder.
func MarshalClusterRegistration(object *ClusterRegistration, target interface{}) error {
	encoder, err := helpers.NewEncoder(target)
	if err != nil {
		return err
	}
	data, err := object.wrap()
	if err != nil {
		return err
	}
	return encoder.Encode(data)
}

// wrap is the method used internally to convert a value of the 'cluster_registration'
// value to a JSON document.
func (o *ClusterRegistration) wrap() (data *clusterRegistrationData, err error) {
	if o == nil {
		return
	}
	data = new(clusterRegistrationData)
	data.SubscriptionID = o.subscriptionID
	data.ExternalID = o.externalID
	return
}

// UnmarshalClusterRegistration reads a value of the 'cluster_registration' type from the given
// source, which can be an slice of bytes, a string, a reader or a JSON decoder.
func UnmarshalClusterRegistration(source interface{}) (object *ClusterRegistration, err error) {
	decoder, err := helpers.NewDecoder(source)
	if err != nil {
		return
	}
	data := new(clusterRegistrationData)
	err = decoder.Decode(data)
	if err != nil {
		return
	}
	object, err = data.unwrap()
	return
}

// unwrap is the function used internally to convert the JSON unmarshalled data to a
// value of the 'cluster_registration' type.
func (d *clusterRegistrationData) unwrap() (object *ClusterRegistration, err error) {
	if d == nil {
		return
	}
	object = new(ClusterRegistration)
	object.subscriptionID = d.SubscriptionID
	object.externalID = d.ExternalID
	return
}
