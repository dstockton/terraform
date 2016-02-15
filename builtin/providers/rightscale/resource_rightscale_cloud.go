package rightscale

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/rightscale/rsc/cm15"
	"github.com/rightscale/rsc/rsapi"
)

func resourceRightScaleCloud() *schema.Resource {
	return &schema.Resource{
		Read: resourceRightScaleCloudRead,

		Schema: map[string]*schema.Schema{
			"cloud_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},

			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},

			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceRightScaleCloudRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cm15.API)

	l := client.CloudLocator("/api/clouds")
	cloud, err := l.Show(rsapi.APIParams{"id": d.Id()})

	if err != nil {
		return fmt.Errorf("Failed to read RightScale cloud: %s", err)
	}

	d.Set("name", cloud.Name)
	d.Set("description", cloud.Description)
	d.Set("cloud_type", cloud.CloudType)

	return nil
}
