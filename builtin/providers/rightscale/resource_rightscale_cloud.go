package rightscale

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/rightscale/rsc/cm15"
	"github.com/rightscale/rsc/rsapi"
)

func resourceRightScaleCloud() *schema.Resource {
	return &schema.Resource{
		Create: resourceRightScaleCloudCreate,
		Read:   resourceRightScaleCloudRead,
		Update: resourceRightScaleCloudUpdate,
		Delete: resourceRightScaleCloudDelete,

		Schema: map[string]*schema.Schema{
			"cloud_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},

			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceRightScaleCloudRead(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[INFO] resourceRightScaleCloudRead")

	client := meta.(*cm15.API)

	l := client.CloudLocator("/api/clouds")
	clouds, err := l.Index(rsapi.APIParams{"filter": "name==" + d.Get("name").(string)})

	if err != nil {
		return fmt.Errorf("Failed to read RightScale cloud: %s", err)
	}

	cloud := clouds[0]

	d.SetId(cloud.Name)

	d.Set("name", cloud.Name)
	d.Set("description", cloud.Description)
	d.Set("cloud_type", cloud.CloudType)

	return nil
}

func resourceRightScaleCloudCreate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[INFO] resourceRightScaleCloudCreate")

	return resourceRightScaleCloudRead(d, meta)
}

func resourceRightScaleCloudUpdate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[INFO] resourceRightScaleCloudUpdate")

	return resourceRightScaleCloudRead(d, meta)
}

func resourceRightScaleCloudDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[INFO] resourceRightScaleCloudDelete")

	return nil
}
