// Copyright 2022 The Sigstore Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package server

import (
	"time"

	"github.com/go-openapi/loads"
	"github.com/navzar05/descentralized-timestamp-authority/pkg/api"
	"github.com/navzar05/descentralized-timestamp-authority/pkg/generated/restapi"
	"github.com/navzar05/descentralized-timestamp-authority/pkg/generated/restapi/operations"
	"github.com/navzar05/descentralized-timestamp-authority/pkg/internal/cmdparams"
)

// NewRestAPIServer creates a server for serving the rest API TSA service
func NewRestAPIServer(host string,
	port int,
	scheme []string,
	httpReadOnly bool,
	readTimeout, writeTimeout time.Duration) *restapi.Server {
	doc, _ := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	server := restapi.NewServer(operations.NewTimestampServerAPI(doc))

	server.Host = host
	server.Port = port
	server.EnabledListeners = scheme
	server.ReadTimeout = readTimeout
	server.WriteTimeout = writeTimeout
	cmdparams.IsHTTPPingOnly = httpReadOnly
	api.ConfigureAPI()
	server.ConfigureAPI()

	return server
}
