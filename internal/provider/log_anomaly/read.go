package loganomaly

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	// "github.com/aws/aws-sdk-go-v2/aws/session"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
	"github.com/skpr/terraform-provider-log-anomaly/internal/provider/log_anomaly/transform"
)

// Read the log anomaly detector.
func Read(d *schema.ResourceData, m interface{}) error {
	cfg := m.(aws.Config)
	c := cloudwatchlogs.NewFromConfig(cfg)

	obj, err := c.GetLogAnomalyDetector(
		context.TODO(),
		&cloudwatchlogs.GetLogAnomalyDetectorInput{
			AnomalyDetectorArn: aws.String(d.Id()),
		},
	)

	if err != nil {
		return errors.Wrap(err, "failed to get ID")
	}

	frequency, err := transform.FromAPIValue(obj.EvaluationFrequency)
	if err != nil {
		return err
	}

	d.Set(Name, obj.DetectorName)
	d.Set(LogGroup, obj.LogGroupArnList)
	d.Set(EvaluationFrequency, frequency)

	return nil
}
