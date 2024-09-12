package loganomaly

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/skpr/terraform-provider-log-anomaly/internal/provider/log_anomaly/transform"
)

// Update the log anomaly detector.
func Update(d *schema.ResourceData, m interface{}) error {
	sess := m.(*session.Session)

	c := cloudwatchlogs.New(sess)

	evaluationFrequency, err := transform.ToAPIValue(d.Get(EvaluationFrequency).(string))

	if err != nil {
		return err
	}

	_, err = c.UpdateLogAnomalyDetector(&cloudwatchlogs.UpdateLogAnomalyDetectorInput{
		AnomalyDetectorArn:  aws.String(d.Id()),
		EvaluationFrequency: aws.String(evaluationFrequency),
		Enabled:             aws.Bool(true),
	})

	if err != nil {
		return err
	}

	return nil
}
