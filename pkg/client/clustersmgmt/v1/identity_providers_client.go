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
	"bytes"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"

	"github.com/openshift-online/uhc-sdk-go/pkg/client/errors"
	"github.com/openshift-online/uhc-sdk-go/pkg/client/helpers"
)

// IdentityProvidersClient is the client of the 'identity_providers' resource.
//
// Manages the collection of identity providers of a cluster.
type IdentityProvidersClient struct {
	transport http.RoundTripper
	path      string
}

// NewIdentityProvidersClient creates a new client for the 'identity_providers'
// resource using the given transport to sned the requests and receive the
// responses.
func NewIdentityProvidersClient(transport http.RoundTripper, path string) *IdentityProvidersClient {
	client := new(IdentityProvidersClient)
	client.transport = transport
	client.path = path
	return client
}

// List creates a request for the 'list' method.
//
// Retrieves the list of identity providers.
func (c *IdentityProvidersClient) List() *IdentityProvidersListRequest {
	request := new(IdentityProvidersListRequest)
	request.transport = c.transport
	request.path = c.path
	return request
}

// Add creates a request for the 'add' method.
//
// Adds a new identity provider to the cluster.
func (c *IdentityProvidersClient) Add() *IdentityProvidersAddRequest {
	request := new(IdentityProvidersAddRequest)
	request.transport = c.transport
	request.path = c.path
	return request
}

// IdentityProvider returns the target 'identity_provider' resource for the given identifier.
//
// Reference to the service that manages an specific identity provider.
func (c *IdentityProvidersClient) IdentityProvider(id string) *IdentityProviderClient {
	return NewIdentityProviderClient(c.transport, path.Join(c.path, id))
}

// IdentityProvidersListRequest is the request for the 'list' method.
type IdentityProvidersListRequest struct {
	transport http.RoundTripper
	path      string
	query     url.Values
	header    http.Header
}

// Parameter adds a query parameter.
func (r *IdentityProvidersListRequest) Parameter(name string, value interface{}) *IdentityProvidersListRequest {
	helpers.AddValue(&r.query, name, value)
	return r
}

// Header adds a request header.
func (r *IdentityProvidersListRequest) Header(name string, value interface{}) *IdentityProvidersListRequest {
	helpers.AddHeader(&r.header, name, value)
	return r
}

// Send sends this request, waits for the response, and returns it.
//
// This is a potentially lengthy operation, as it requires network communication.
// Consider using a context and the SendContext method. If you don't provide a
// context then a new background context will be created.
func (r *IdentityProvidersListRequest) Send() (result *IdentityProvidersListResponse, err error) {
	return r.SendContext(context.Background())
}

// SendContext sends this request, waits for the response, and returns it.
func (r *IdentityProvidersListRequest) SendContext(ctx context.Context) (result *IdentityProvidersListResponse, err error) {
	query := helpers.CopyQuery(r.query)
	header := helpers.CopyHeader(r.header)
	uri := &url.URL{
		Path:     r.path,
		RawQuery: query.Encode(),
	}
	request := &http.Request{
		Method: http.MethodGet,
		URL:    uri,
		Header: header,
	}
	if ctx != nil {
		request = request.WithContext(ctx)
	}
	response, err := r.transport.RoundTrip(request)
	if err != nil {
		return
	}
	defer response.Body.Close()
	result = new(IdentityProvidersListResponse)
	result.status = response.StatusCode
	result.header = response.Header
	if result.status >= 400 {
		result.err, err = errors.UnmarshalError(response.Body)
		if err != nil {
			return
		}
		err = result.err
		return
	}
	err = result.unmarshal(response.Body)
	if err != nil {
		return
	}
	return
}

// IdentityProvidersListResponse is the response for the 'list' method.
type IdentityProvidersListResponse struct {
	status int
	header http.Header
	err    *errors.Error
	page   *int
	size   *int
	total  *int
	items  *IdentityProviderList
}

// Status returns the response status code.
func (r *IdentityProvidersListResponse) Status() int {
	return r.status
}

// Header returns header of the response.
func (r *IdentityProvidersListResponse) Header() http.Header {
	return r.header
}

// Error returns the response error.
func (r *IdentityProvidersListResponse) Error() *errors.Error {
	return r.err
}

// Page returns the value of the 'page' parameter.
//
// Index of the requested page, where one corresponds to the first page.
func (r *IdentityProvidersListResponse) Page() int {
	if r.page != nil {
		return *r.page
	}
	return 0
}

// Size returns the value of the 'size' parameter.
//
// Number of items contained in the returned page.
func (r *IdentityProvidersListResponse) Size() int {
	if r.size != nil {
		return *r.size
	}
	return 0
}

// Total returns the value of the 'total' parameter.
//
// Total number of items of the collection.
func (r *IdentityProvidersListResponse) Total() int {
	if r.total != nil {
		return *r.total
	}
	return 0
}

// Items returns the value of the 'items' parameter.
//
// Retrieved list of identity providers.
func (r *IdentityProvidersListResponse) Items() *IdentityProviderList {
	return r.items
}

// unmarshal is the method used internally to unmarshal responses to the
// 'list' method.
func (r *IdentityProvidersListResponse) unmarshal(reader io.Reader) error {
	var err error
	decoder := json.NewDecoder(reader)
	data := new(identityProvidersListResponseData)
	err = decoder.Decode(data)
	if err != nil {
		return err
	}
	r.page = data.Page
	r.size = data.Size
	r.total = data.Total
	r.items, err = data.Items.unwrap()
	if err != nil {
		return err
	}
	return err
}

// identityProvidersListResponseData is the structure used internally to unmarshal
// the response of the 'list' method.
type identityProvidersListResponseData struct {
	Page  *int                     "json:\"page,omitempty\""
	Size  *int                     "json:\"size,omitempty\""
	Total *int                     "json:\"total,omitempty\""
	Items identityProviderListData "json:\"items,omitempty\""
}

// IdentityProvidersAddRequest is the request for the 'add' method.
type IdentityProvidersAddRequest struct {
	transport http.RoundTripper
	path      string
	query     url.Values
	header    http.Header
	body      *IdentityProvider
}

// Parameter adds a query parameter.
func (r *IdentityProvidersAddRequest) Parameter(name string, value interface{}) *IdentityProvidersAddRequest {
	helpers.AddValue(&r.query, name, value)
	return r
}

// Header adds a request header.
func (r *IdentityProvidersAddRequest) Header(name string, value interface{}) *IdentityProvidersAddRequest {
	helpers.AddHeader(&r.header, name, value)
	return r
}

// Body sets the value of the 'body' parameter.
//
// Description of the cluster.
func (r *IdentityProvidersAddRequest) Body(value *IdentityProvider) *IdentityProvidersAddRequest {
	r.body = value
	return r
}

// Send sends this request, waits for the response, and returns it.
//
// This is a potentially lengthy operation, as it requires network communication.
// Consider using a context and the SendContext method. If you don't provide a
// context then a new background context will be created.
func (r *IdentityProvidersAddRequest) Send() (result *IdentityProvidersAddResponse, err error) {
	return r.SendContext(context.Background())
}

// SendContext sends this request, waits for the response, and returns it.
func (r *IdentityProvidersAddRequest) SendContext(ctx context.Context) (result *IdentityProvidersAddResponse, err error) {
	query := helpers.CopyQuery(r.query)
	header := helpers.CopyHeader(r.header)
	buffer := new(bytes.Buffer)
	err = r.marshal(buffer)
	if err != nil {
		return
	}
	uri := &url.URL{
		Path:     r.path,
		RawQuery: query.Encode(),
	}
	request := &http.Request{
		Method: http.MethodPost,
		URL:    uri,
		Header: header,
		Body:   ioutil.NopCloser(buffer),
	}
	if ctx != nil {
		request = request.WithContext(ctx)
	}
	response, err := r.transport.RoundTrip(request)
	if err != nil {
		return
	}
	defer response.Body.Close()
	result = new(IdentityProvidersAddResponse)
	result.status = response.StatusCode
	result.header = response.Header
	if result.status >= 400 {
		result.err, err = errors.UnmarshalError(response.Body)
		if err != nil {
			return
		}
		err = result.err
		return
	}
	err = result.unmarshal(response.Body)
	if err != nil {
		return
	}
	return
}

// marshall is the method used internally to marshal requests for the
// 'add' method.
func (r *IdentityProvidersAddRequest) marshal(writer io.Writer) error {
	var err error
	encoder := json.NewEncoder(writer)
	data, err := r.body.wrap()
	if err != nil {
		return err
	}
	err = encoder.Encode(data)
	return err
}

// IdentityProvidersAddResponse is the response for the 'add' method.
type IdentityProvidersAddResponse struct {
	status int
	header http.Header
	err    *errors.Error
	body   *IdentityProvider
}

// Status returns the response status code.
func (r *IdentityProvidersAddResponse) Status() int {
	return r.status
}

// Header returns header of the response.
func (r *IdentityProvidersAddResponse) Header() http.Header {
	return r.header
}

// Error returns the response error.
func (r *IdentityProvidersAddResponse) Error() *errors.Error {
	return r.err
}

// Body returns the value of the 'body' parameter.
//
// Description of the cluster.
func (r *IdentityProvidersAddResponse) Body() *IdentityProvider {
	return r.body
}

// unmarshal is the method used internally to unmarshal responses to the
// 'add' method.
func (r *IdentityProvidersAddResponse) unmarshal(reader io.Reader) error {
	var err error
	decoder := json.NewDecoder(reader)
	data := new(identityProviderData)
	err = decoder.Decode(data)
	if err != nil {
		return err
	}
	r.body, err = data.unwrap()
	if err != nil {
		return err
	}
	return err
}
