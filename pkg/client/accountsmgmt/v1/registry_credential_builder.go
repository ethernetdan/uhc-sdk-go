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

package v1 // github.com/openshift-online/uhc-sdk-go/pkg/client/accountsmgmt/v1

// RegistryCredentialBuilder contains the data and logic needed to build 'registry_credential' objects.
//
//
type RegistryCredentialBuilder struct {
	id       *string
	href     *string
	link     bool
	username *string
	token    *string
	registry *RegistryBuilder
	account  *AccountBuilder
}

// NewRegistryCredential creates a new builder of 'registry_credential' objects.
func NewRegistryCredential() *RegistryCredentialBuilder {
	return new(RegistryCredentialBuilder)
}

// ID sets the identifier of the object.
func (b *RegistryCredentialBuilder) ID(value string) *RegistryCredentialBuilder {
	b.id = &value
	return b
}

// HREF sets the link to the object.
func (b *RegistryCredentialBuilder) HREF(value string) *RegistryCredentialBuilder {
	b.href = &value
	return b
}

// Link sets the flag that indicates if this is a link.
func (b *RegistryCredentialBuilder) Link(value bool) *RegistryCredentialBuilder {
	b.link = value
	return b
}

// Username sets the value of the 'username' attribute
// to the given value.
//
//
func (b *RegistryCredentialBuilder) Username(value string) *RegistryCredentialBuilder {
	b.username = &value
	return b
}

// Token sets the value of the 'token' attribute
// to the given value.
//
//
func (b *RegistryCredentialBuilder) Token(value string) *RegistryCredentialBuilder {
	b.token = &value
	return b
}

// Registry sets the value of the 'registry' attribute
// to the given value.
//
//
func (b *RegistryCredentialBuilder) Registry(value *RegistryBuilder) *RegistryCredentialBuilder {
	b.registry = value
	return b
}

// Account sets the value of the 'account' attribute
// to the given value.
//
//
func (b *RegistryCredentialBuilder) Account(value *AccountBuilder) *RegistryCredentialBuilder {
	b.account = value
	return b
}

// Build creates a 'registry_credential' object using the configuration stored in the builder.
func (b *RegistryCredentialBuilder) Build() (object *RegistryCredential, err error) {
	object = new(RegistryCredential)
	object.id = b.id
	object.href = b.href
	object.link = b.link
	if b.username != nil {
		object.username = b.username
	}
	if b.token != nil {
		object.token = b.token
	}
	if b.registry != nil {
		object.registry, err = b.registry.Build()
		if err != nil {
			return
		}
	}
	if b.account != nil {
		object.account, err = b.account.Build()
		if err != nil {
			return
		}
	}
	return
}
