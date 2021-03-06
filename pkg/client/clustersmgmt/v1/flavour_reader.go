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
	"fmt"

	"github.com/openshift-online/uhc-sdk-go/pkg/client/helpers"
)

// flavourData is the data structure used internally to marshal and unmarshal
// objects of type 'flavour'.
type flavourData struct {
	Kind    *string           "json:\"kind,omitempty\""
	ID      *string           "json:\"id,omitempty\""
	HREF    *string           "json:\"href,omitempty\""
	AWS     *awsData          "json:\"aws,omitempty\""
	Version *string           "json:\"version,omitempty\""
	Nodes   *clusterNodesData "json:\"nodes,omitempty\""
	Name    *string           "json:\"name,omitempty\""
	Network *networkData      "json:\"network,omitempty\""
}

// MarshalFlavour writes a value of the 'flavour' to the given target,
// which can be a writer or a JSON encoder.
func MarshalFlavour(object *Flavour, target interface{}) error {
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

// wrap is the method used internally to convert a value of the 'flavour'
// value to a JSON document.
func (o *Flavour) wrap() (data *flavourData, err error) {
	if o == nil {
		return
	}
	data = new(flavourData)
	data.ID = o.id
	data.HREF = o.href
	data.Kind = new(string)
	if o.link {
		*data.Kind = FlavourLinkKind
	} else {
		*data.Kind = FlavourKind
	}
	data.AWS, err = o.aws.wrap()
	if err != nil {
		return
	}
	data.Version = o.version
	data.Nodes, err = o.nodes.wrap()
	if err != nil {
		return
	}
	data.Name = o.name
	data.Network, err = o.network.wrap()
	if err != nil {
		return
	}
	return
}

// UnmarshalFlavour reads a value of the 'flavour' type from the given
// source, which can be an slice of bytes, a string, a reader or a JSON decoder.
func UnmarshalFlavour(source interface{}) (object *Flavour, err error) {
	decoder, err := helpers.NewDecoder(source)
	if err != nil {
		return
	}
	data := new(flavourData)
	err = decoder.Decode(data)
	if err != nil {
		return
	}
	object, err = data.unwrap()
	return
}

// unwrap is the function used internally to convert the JSON unmarshalled data to a
// value of the 'flavour' type.
func (d *flavourData) unwrap() (object *Flavour, err error) {
	if d == nil {
		return
	}
	object = new(Flavour)
	object.id = d.ID
	object.href = d.HREF
	if d.Kind != nil {
		switch *d.Kind {
		case FlavourKind:
			object.link = false
		case FlavourLinkKind:
			object.link = true
		default:
			err = fmt.Errorf(
				"expected kind '%s' or '%s' but got '%s'",
				FlavourKind,
				FlavourLinkKind,
				*d.Kind,
			)
			return
		}
	}
	object.aws, err = d.AWS.unwrap()
	if err != nil {
		return
	}
	object.version = d.Version
	object.nodes, err = d.Nodes.unwrap()
	if err != nil {
		return
	}
	object.name = d.Name
	object.network, err = d.Network.unwrap()
	if err != nil {
		return
	}
	return
}
