package rightscale

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

// Provider returns a terraform.ResourceProvider.
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"refreshToken": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("RIGHTSCALE_REFRESHTOKEN", nil),
				Description: "RightScale OAuth refresh token.",
			},

			"accountID": &schema.Schema{
				Type:        schema.TypeInt,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("RIGHTSCALE_ACCOUNTID", nil),
				Description: "RightScale account ID to work with.",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"rightscale_cloud": resourceRightScaleCloud(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		RefreshToken: d.Get("refreshToken").(string),
		AccountID:    d.Get("accountID").(int),
	}

	return config.Client()
}
