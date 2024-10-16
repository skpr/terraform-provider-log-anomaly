package loggroup

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	// Name of the anomaly detector
	Name = "name"
	// LogGroup ARN that needs anomaly detection setup for.
	RetentionInDays = "retention_in_days"
	// EvaluationFrequency in minutes for the anomaly detector.
	EvaluationFrequency = "evaluation_frequency"
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
			RetentionInDays: {
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	}
}
