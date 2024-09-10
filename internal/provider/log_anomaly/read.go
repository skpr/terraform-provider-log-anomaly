package loganomaly

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pkg/errors"
)

// Read the S3 object.
func Read(d *schema.ResourceData, m interface{}) error {
	c := cloudwatchlogs.New(m.(*session.Session))

	obj, err := c.GetLogAnomalyDetector(&cloudwatchlogs.GetLogAnomalyDetectorInput{
		AnomalyDetectorArn: aws.String(d.Id()),
	})

	if err != nil {
		return errors.Wrap(err, "failed to get ID")
	}

	d.Set(Name, obj.DetectorName)
	d.Set(LogGroup, obj.LogGroupArnList)
	d.Set(EvaluationFrequency, obj.EvaluationFrequency)

	return nil
}
