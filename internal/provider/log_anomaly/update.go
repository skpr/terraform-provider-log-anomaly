package loganomaly

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	// "github.com/aws/aws-sdk-go-v2/aws/session"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/skpr/terraform-provider-log-anomaly/internal/provider/log_anomaly/transform"
)

// Update the log anomaly detector.
func Update(d *schema.ResourceData, m interface{}) error {
	cfg := m.(aws.Config)
	c := cloudwatchlogs.NewFromConfig(cfg)

	evaluationFrequency, err := transform.ToAPIValue(d.Get(EvaluationFrequency).(string))

	if err != nil {
		return err
	}

	_, err = c.UpdateLogAnomalyDetector(
		context.TODO(),
		&cloudwatchlogs.UpdateLogAnomalyDetectorInput{
			AnomalyDetectorArn:  aws.String(d.Id()),
			EvaluationFrequency: evaluationFrequency,
			Enabled:             aws.Bool(true),
		},
	)

	if err != nil {
		return err
	}

	return nil
}
