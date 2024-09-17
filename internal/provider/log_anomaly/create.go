package loganomaly

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/skpr/terraform-provider-log-anomaly/internal/provider/log_anomaly/transform"
)

// Create the log anomaly detector.
func Create(d *schema.ResourceData, m interface{}) error {
	cfg := m.(aws.Config)
	c := cloudwatchlogs.NewFromConfig(cfg)

	var (
		name     = d.Get(Name).(string)
		logGroup = d.Get(LogGroup).(string)
	)

	evaluationFrequency, err := transform.ToAPIValue(d.Get(EvaluationFrequency).(string))

	if err != nil {
		return err
	}

	obj, err := c.CreateLogAnomalyDetector(
		context.TODO(),
		&cloudwatchlogs.CreateLogAnomalyDetectorInput{
			DetectorName:        aws.String(name),
			LogGroupArnList:     []string{logGroup},
			EvaluationFrequency: evaluationFrequency,
		},
	)

	if err != nil {
		return err
	}

	d.SetId(*obj.AnomalyDetectorArn)

	return nil
}
