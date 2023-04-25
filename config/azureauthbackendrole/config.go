package DIR_azureauthbackendrole

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vault_azure_auth_backend_role", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "vault"
	})
}
