package provider


type HttpProviderConfig struct {
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/http/3.3.0/docs#alias HttpProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
}

