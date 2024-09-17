package loganomaly

import (
	// "github.com/aws/aws-sdk-go-v2/aws/session"

	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Delete the log anomaly detector.
func Delete(d *schema.ResourceData, m interface{}) error {
	cfg := m.(aws.Config)
	c := cloudwatchlogs.NewFromConfig(cfg)

	_, err := c.DeleteLogAnomalyDetector(
		context.TODO(),
		&cloudwatchlogs.DeleteLogAnomalyDetectorInput{
			AnomalyDetectorArn: &d.State().ID,
		},
	)
	if err != nil {
		return err
	}

	return nil
}
