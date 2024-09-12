package loganomaly

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/skpr/terraform-provider-log-anomaly/internal/provider/log_anomaly/transform"
)

// Create the log anomaly detector.
func Create(d *schema.ResourceData, m interface{}) error {
	sess := m.(*session.Session)

	var (
		name     = d.Get(Name).(string)
		logGroup = d.Get(LogGroup).(string)
	)

	evaluationFrequency, err := transform.ToAPIValue(d.Get(EvaluationFrequency).(string))

	if err != nil {
		return err
	}

	c := cloudwatchlogs.New(sess)

	obj, err := c.CreateLogAnomalyDetector(&cloudwatchlogs.CreateLogAnomalyDetectorInput{
		DetectorName:        aws.String(name),
		LogGroupArnList:     aws.StringSlice([]string{logGroup}),
		EvaluationFrequency: aws.String(evaluationFrequency),
	})

	if err != nil {
		return err
	}

	d.SetId(*obj.AnomalyDetectorArn)

	return nil
}
