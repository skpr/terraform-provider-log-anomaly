package loganomaly

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/hashicorp/terraform/helper/schema"
)

// Delete the S3 object.
func Delete(d *schema.ResourceData, m interface{}) error {
	sess := m.(*session.Session)

	c := cloudwatchlogs.New(sess)

	_, err := c.DeleteLogAnomalyDetector(&cloudwatchlogs.DeleteLogAnomalyDetectorInput{
		AnomalyDetectorArn: &d.State().ID,
	})
	if err != nil {
		return err
	}

	return nil
}
