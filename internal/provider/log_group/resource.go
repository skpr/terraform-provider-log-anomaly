package loggroup

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	// Name of the log group
	Name = "name"
)

// Resource returns this packages resource.
func Resource() *schema.Resource {
	return &schema.Resource{
		Create: Create,
		Read:   Read,
		Update: Update,
		Delete: Delete,

		Schema: map[string]*schema.Schema{
			Name: {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}
