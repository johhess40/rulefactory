package cmd

type AdoAuth struct {
	ClientId     string `mapstructure:"client_id"`
	ClientSecret string `mapstructure:"client_secret"`
	TenantId     string `mapstructure:"tenant_id"`
}
