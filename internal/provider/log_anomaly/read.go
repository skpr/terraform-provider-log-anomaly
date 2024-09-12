package loganomaly

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pkg/errors"
	"github.com/skpr/terraform-provider-log-anomaly/internal/provider/log_anomaly/transform"
)

// Read the log anomaly detector.
func Read(d *schema.ResourceData, m interface{}) error {
	c := cloudwatchlogs.New(m.(*session.Session))

	obj, err := c.GetLogAnomalyDetector(&cloudwatchlogs.GetLogAnomalyDetectorInput{
		AnomalyDetectorArn: aws.String(d.Id()),
	})

	if err != nil {
		return errors.Wrap(err, "failed to get ID")
	}

	frequency, err := transform.FromApiValue(*obj.EvaluationFrequency)
	if err != nil {
		return err
	}

	d.Set(Name, obj.DetectorName)
	d.Set(LogGroup, obj.LogGroupArnList)
	d.Set(EvaluationFrequency, frequency)

	return nil
}
