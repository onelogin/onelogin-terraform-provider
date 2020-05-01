package sso

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// AppConfiguration returns a key/value map of the various fields that make up
// the AppConfiguration field for a OneLogin App.
func OIDCSSOSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"client_id": &schema.Schema{
			Type:     schema.TypeString,
			Computed: true,
		},
		"client_secret": &schema.Schema{
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}

// AppConfiguration returns a key/value map of the various fields that make up
// the AppConfiguration field for a OneLogin App.
func SAMLSSOSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"acs_url": &schema.Schema{
			Type:     schema.TypeString,
			Computed: true,
		},
		"metadata_url": &schema.Schema{
			Type:     schema.TypeString,
			Computed: true,
		},
		"issuer": &schema.Schema{
			Type:     schema.TypeString,
			Computed: true,
		},
		"certificate": &schema.Schema{
			Type:     schema.TypeSet,
			MaxItems: 1,
			Computed: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"name": &schema.Schema{
						Type:     schema.TypeString,
						Computed: true,
					},
					"id": &schema.Schema{
						Type:     schema.TypeInt,
						Computed: true,
					},
					"value": &schema.Schema{
						Type:     schema.TypeString,
						Computed: true,
					},
					"sls_url": &schema.Schema{
						Type:     schema.TypeString,
						Computed: true,
					},
				},
			},
		},
	}
}