package datahttp

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-http.dataHttp.DataHttp",
		reflect.TypeOf((*DataHttp)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "body", GoGetter: "Body"},
			_jsii_.MemberProperty{JsiiProperty: "caCertPem", GoGetter: "CaCertPem"},
			_jsii_.MemberProperty{JsiiProperty: "caCertPemInput", GoGetter: "CaCertPemInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "insecure", GoGetter: "Insecure"},
			_jsii_.MemberProperty{JsiiProperty: "insecureInput", GoGetter: "InsecureInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "method", GoGetter: "Method"},
			_jsii_.MemberProperty{JsiiProperty: "methodInput", GoGetter: "MethodInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberMethod{JsiiMethod: "putRetry", GoMethod: "PutRetry"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "requestBody", GoGetter: "RequestBody"},
			_jsii_.MemberProperty{JsiiProperty: "requestBodyInput", GoGetter: "RequestBodyInput"},
			_jsii_.MemberProperty{JsiiProperty: "requestHeaders", GoGetter: "RequestHeaders"},
			_jsii_.MemberProperty{JsiiProperty: "requestHeadersInput", GoGetter: "RequestHeadersInput"},
			_jsii_.MemberProperty{JsiiProperty: "requestTimeoutMs", GoGetter: "RequestTimeoutMs"},
			_jsii_.MemberProperty{JsiiProperty: "requestTimeoutMsInput", GoGetter: "RequestTimeoutMsInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetCaCertPem", GoMethod: "ResetCaCertPem"},
			_jsii_.MemberMethod{JsiiMethod: "resetInsecure", GoMethod: "ResetInsecure"},
			_jsii_.MemberMethod{JsiiMethod: "resetMethod", GoMethod: "ResetMethod"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetRequestBody", GoMethod: "ResetRequestBody"},
			_jsii_.MemberMethod{JsiiMethod: "resetRequestHeaders", GoMethod: "ResetRequestHeaders"},
			_jsii_.MemberMethod{JsiiMethod: "resetRequestTimeoutMs", GoMethod: "ResetRequestTimeoutMs"},
			_jsii_.MemberMethod{JsiiMethod: "resetRetry", GoMethod: "ResetRetry"},
			_jsii_.MemberProperty{JsiiProperty: "responseBody", GoGetter: "ResponseBody"},
			_jsii_.MemberProperty{JsiiProperty: "responseBodyBase64", GoGetter: "ResponseBodyBase64"},
			_jsii_.MemberProperty{JsiiProperty: "responseHeaders", GoGetter: "ResponseHeaders"},
			_jsii_.MemberProperty{JsiiProperty: "retry", GoGetter: "Retry"},
			_jsii_.MemberProperty{JsiiProperty: "retryInput", GoGetter: "RetryInput"},
			_jsii_.MemberProperty{JsiiProperty: "statusCode", GoGetter: "StatusCode"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "url", GoGetter: "Url"},
			_jsii_.MemberProperty{JsiiProperty: "urlInput", GoGetter: "UrlInput"},
		},
		func() interface{} {
			j := jsiiProxy_DataHttp{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformDataSource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-http.dataHttp.DataHttpConfig",
		reflect.TypeOf((*DataHttpConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-http.dataHttp.DataHttpRetry",
		reflect.TypeOf((*DataHttpRetry)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-http.dataHttp.DataHttpRetryOutputReference",
		reflect.TypeOf((*DataHttpRetryOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "attempts", GoGetter: "Attempts"},
			_jsii_.MemberProperty{JsiiProperty: "attemptsInput", GoGetter: "AttemptsInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "maxDelayMs", GoGetter: "MaxDelayMs"},
			_jsii_.MemberProperty{JsiiProperty: "maxDelayMsInput", GoGetter: "MaxDelayMsInput"},
			_jsii_.MemberProperty{JsiiProperty: "minDelayMs", GoGetter: "MinDelayMs"},
			_jsii_.MemberProperty{JsiiProperty: "minDelayMsInput", GoGetter: "MinDelayMsInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAttempts", GoMethod: "ResetAttempts"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxDelayMs", GoMethod: "ResetMaxDelayMs"},
			_jsii_.MemberMethod{JsiiMethod: "resetMinDelayMs", GoMethod: "ResetMinDelayMs"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataHttpRetryOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
