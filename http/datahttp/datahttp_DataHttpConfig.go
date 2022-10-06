package datahttp

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataHttpConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/http/d/http#url DataHttp#url}
	Url *string `field:"required" json:"url" yaml:"url"`
	// The HTTP Method for the request.
	//
	// Allowed methods are a subset of methods defined in [RFC7231](https://datatracker.ietf.org/doc/html/rfc7231#section-4.3) namely, `GET`, `HEAD`, and `POST`. `POST` support is only intended for read-only URLs, such as submitting a search.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/http/d/http#method DataHttp#method}
	Method *string `field:"optional" json:"method" yaml:"method"`
	// The request body as a string.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/http/d/http#request_body DataHttp#request_body}
	RequestBody *string `field:"optional" json:"requestBody" yaml:"requestBody"`
	// A map of request header field names and values.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/http/d/http#request_headers DataHttp#request_headers}
	RequestHeaders *map[string]*string `field:"optional" json:"requestHeaders" yaml:"requestHeaders"`
}

