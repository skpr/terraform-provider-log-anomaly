package loganomaly

import (
	"github.com/hashicorp/terraform/helper/schema"
)

const (
	// Name of the anomaly detector
	Name = "name"
	// LogGroup ARN that needs anomaly detection setup for.
	LogGroup = "log_group"
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
			LogGroup: {
				Type:     schema.TypeString,
				Required: true,
			},
			EvaluationFrequency: {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}
