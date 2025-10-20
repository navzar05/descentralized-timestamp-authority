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

package client

import (
	"net/url"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/navzar05/descentralized-timestamp-authority/pkg/generated/client"
)

func GetTimestampClient(timestampServerURL string, opts ...Option) (*client.TimestampAuthority, error) {
	url, err := url.Parse(timestampServerURL)
	if err != nil {
		return nil, err
	}
	o := makeOptions(opts...)

	rt := httptransport.New(url.Host, client.DefaultBasePath, []string{url.Scheme})
	// Input to server
	rt.Producers["application/timestamp-query"] = runtime.ByteStreamProducer()
	rt.Producers["application/json"] = runtime.JSONProducer()
	// Output from server
	rt.Consumers["application/timestamp-reply"] = runtime.ByteStreamConsumer()
	rt.Consumers["application/json"] = runtime.JSONConsumer()
	rt.Consumers["application/pem-certificate-chain"] = runtime.TextConsumer()

	rt.Transport = createRoundTripper(rt.Transport, o)

	registry := strfmt.Default
	return client.New(rt, registry), nil
}
