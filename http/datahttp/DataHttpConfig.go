// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datahttp

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataHttpConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The URL for the request. Supported schemes are `http` and `https`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/http/3.4.1/docs/data-sources/http#url DataHttp#url}
	Url *string `field:"required" json:"url" yaml:"url"`
	// Certificate data of the Certificate Authority (CA) in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/http/3.4.1/docs/data-sources/http#ca_cert_pem DataHttp#ca_cert_pem}
	CaCertPem *string `field:"optional" json:"caCertPem" yaml:"caCertPem"`
	// Disables verification of the server's certificate chain and hostname. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/http/3.4.1/docs/data-sources/http#insecure DataHttp#insecure}
	Insecure interface{} `field:"optional" json:"insecure" yaml:"insecure"`
	// The HTTP Method for the request.
	//
	// Allowed methods are a subset of methods defined in [RFC7231](https://datatracker.ietf.org/doc/html/rfc7231#section-4.3) namely, `GET`, `HEAD`, and `POST`. `POST` support is only intended for read-only URLs, such as submitting a search.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/http/3.4.1/docs/data-sources/http#method DataHttp#method}
	Method *string `field:"optional" json:"method" yaml:"method"`
	// The request body as a string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/http/3.4.1/docs/data-sources/http#request_body DataHttp#request_body}
	RequestBody *string `field:"optional" json:"requestBody" yaml:"requestBody"`
	// A map of request header field names and values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/http/3.4.1/docs/data-sources/http#request_headers DataHttp#request_headers}
	RequestHeaders *map[string]*string `field:"optional" json:"requestHeaders" yaml:"requestHeaders"`
	// The request timeout in milliseconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/http/3.4.1/docs/data-sources/http#request_timeout_ms DataHttp#request_timeout_ms}
	RequestTimeoutMs *float64 `field:"optional" json:"requestTimeoutMs" yaml:"requestTimeoutMs"`
	// retry block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/http/3.4.1/docs/data-sources/http#retry DataHttp#retry}
	Retry *DataHttpRetry `field:"optional" json:"retry" yaml:"retry"`
}

