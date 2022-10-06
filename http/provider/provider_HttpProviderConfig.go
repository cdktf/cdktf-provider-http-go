package provider


type HttpProviderConfig struct {
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/http#alias HttpProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
}

