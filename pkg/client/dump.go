/*
Copyright (c) 2018 Red Hat, Inc.

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

// This file contains the implementations of the methods of the connection that are used to dump to
// the log the details of HTTP requests and responses.

package client

import (
	"context"
	"encoding/json"
	"net/http"
	"sort"
	"strings"
)

// dumpRequest dumps to the log, in debug level, the details of the given HTTP request.
func (c *Connection) dumpRequest(ctx context.Context, request *http.Request, body []byte) {
	c.logger.Debug(ctx, "Request method is %s", request.Method)
	c.logger.Debug(ctx, "Request URL is '%s'", request.URL)
	header := request.Header
	names := make([]string, len(header))
	i := 0
	for name := range header {
		names[i] = name
		i++
	}
	sort.Strings(names)
	for _, name := range names {
		values := header[name]
		for _, value := range values {
			if strings.ToLower(name) == "authorization" {
				c.logger.Debug(ctx, "Request header '%s' is omitted", name)
			} else {
				c.logger.Debug(ctx, "Request header '%s' is '%s'", name, value)
			}
		}
	}
	if body != nil {
		c.logger.Debug(ctx, "Request body follows")
		c.dumpBody(ctx, header, body)
	}
}

// dumpResponse dumps to the log, in debug level, the details of the given HTTP response.
func (c *Connection) dumpResponse(ctx context.Context, response *http.Response, body []byte) {
	c.logger.Debug(ctx, "Response status is '%s'", response.Status)
	c.logger.Debug(ctx, "Response status code %d", response.StatusCode)
	header := response.Header
	names := make([]string, len(header))
	i := 0
	for name := range header {
		names[i] = name
		i++
	}
	sort.Strings(names)
	for _, name := range names {
		values := header[name]
		for _, value := range values {
			c.logger.Debug(ctx, "Response header '%s' is '%s'", name, value)
		}
	}
	if body != nil {
		c.logger.Debug(ctx, "Response body follows")
		c.dumpBody(ctx, header, body)
	}
}

// dumpBody checks the content type used in the given header and then it dumps the given body in a
// format suitable for that content type.
func (c *Connection) dumpBody(ctx context.Context, header http.Header, body []byte) {
	switch header.Get("Content-Type") {
	case "application/json", "":
		c.dumpJSON(ctx, body)
	default:
		c.dumpBytes(ctx, body)
	}
}

// dumpJSON tries to parse the given data as a JSON document. If that works, then it dumps it
// indented, otherwise dumps it as is.
func (c *Connection) dumpJSON(ctx context.Context, data []byte) {
	var parsed map[string]interface{}
	err := json.Unmarshal(data, &parsed)
	if err != nil {
		c.logger.Debug(ctx, "%s", data)
	} else {
		indented, err := json.MarshalIndent(parsed, "", "  ")
		if err != nil {
			c.logger.Debug(ctx, "%s", data)
		} else {
			c.logger.Debug(ctx, "%s", indented)
		}
	}
}

// dumpBytes dump the given data as an array of bytes.
func (c *Connection) dumpBytes(ctx context.Context, data []byte) {
	c.logger.Debug(ctx, "%s", data)
}
