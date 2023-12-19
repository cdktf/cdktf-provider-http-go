// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datahttp


type DataHttpRetry struct {
	// The number of times the request is to be retried.
	//
	// For example, if 2 is specified, the request will be tried a maximum of 3 times.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/http/3.4.1/docs/data-sources/http#attempts DataHttp#attempts}
	Attempts *float64 `field:"optional" json:"attempts" yaml:"attempts"`
	// The maximum delay between retry requests in milliseconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/http/3.4.1/docs/data-sources/http#max_delay_ms DataHttp#max_delay_ms}
	MaxDelayMs *float64 `field:"optional" json:"maxDelayMs" yaml:"maxDelayMs"`
	// The minimum delay between retry requests in milliseconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/http/3.4.1/docs/data-sources/http#min_delay_ms DataHttp#min_delay_ms}
	MinDelayMs *float64 `field:"optional" json:"minDelayMs" yaml:"minDelayMs"`
}

