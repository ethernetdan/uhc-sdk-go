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

// userListData is type used internally to marshal and unmarshal lists of objects
// of type 'user'.
type userListData []*userData

// UnmarshalUserList reads a list of values of the 'user'
// from the given source, which can be a slice of bytes, a string, an io.Reader or a
// json.Decoder.
func UnmarshalUserList(source interface{}) (list *UserList, err error) {
	decoder, err := helpers.NewDecoder(source)
	if err != nil {
		return
	}
	var data userListData
	err = decoder.Decode(&data)
	if err != nil {
		return
	}
	list, err = data.unwrap()
	return
}

// wrap is the method used internally to convert a list of values of the
// 'user' value to a JSON document.
func (l *UserList) wrap() (data userListData, err error) {
	if l == nil {
		return
	}
	data = make(userListData, len(l.items))
	for i, item := range l.items {
		data[i], err = item.wrap()
		if err != nil {
			return
		}
	}
	return
}

// unwrap is the function used internally to convert the JSON unmarshalled data to a
// list of values of the 'user' type.
func (d userListData) unwrap() (list *UserList, err error) {
	if d == nil {
		return
	}
	items := make([]*User, len(d))
	for i, item := range d {
		items[i], err = item.unwrap()
		if err != nil {
			return
		}
	}
	list = new(UserList)
	list.items = items
	return
}

// userListLinkData is type used internally to marshal and unmarshal links
// to lists of objects of type 'user'.
type userListLinkData struct {
	Kind  *string     "json:\"kind,omitempty\""
	HREF  *string     "json:\"href,omitempty\""
	Items []*userData "json:\"items,omitempty\""
}

// wrapLink is the method used internally to convert a list of values of the
// 'user' value to a link.
func (l *UserList) wrapLink() (data *userListLinkData, err error) {
	if l == nil {
		return
	}
	items := make([]*userData, len(l.items))
	for i, item := range l.items {
		items[i], err = item.wrap()
		if err != nil {
			return
		}
	}
	data = new(userListLinkData)
	data.Items = items
	data.HREF = l.href
	data.Kind = new(string)
	if l.link {
		*data.Kind = UserListLinkKind
	} else {
		*data.Kind = UserListKind
	}
	return
}

// unwrapLink is the function used internally to convert a JSON link to a list
// of values of the 'user' type to a list.
func (d *userListLinkData) unwrapLink() (list *UserList, err error) {
	if d == nil {
		return
	}
	items := make([]*User, len(d.Items))
	for i, item := range d.Items {
		items[i], err = item.unwrap()
		if err != nil {
			return
		}
	}
	list = new(UserList)
	list.items = items
	list.href = d.HREF
	if d.Kind != nil {
		switch *d.Kind {
		case UserListKind:
			list.link = false
		case UserListLinkKind:
			list.link = true
		default:
			err = fmt.Errorf(
				"expected kind '%s' or '%s' but got '%s'",
				UserListKind,
				UserListLinkKind,
				*d.Kind,
			)
			return
		}
	}
	return
}
