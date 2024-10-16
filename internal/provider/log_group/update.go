package loggroup

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	// "github.com/aws/aws-sdk-go-v2/aws/session"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Update the log anomaly detector.
func Update(d *schema.ResourceData, m interface{}) error {
	cfg := m.(aws.Config)
	c := cloudwatchlogs.NewFromConfig(cfg)

	var rid int
	if d.Get(RetentionInDays).(int) < 1 {
		rid = 90
	} else {
		rid = d.Get(RetentionInDays).(int)
	}

	logGroupName := d.Get("name").(string)

	if d.HasChange(RetentionInDays) {
		_, err := c.PutRetentionPolicy(
			context.TODO(),
			&cloudwatchlogs.PutRetentionPolicyInput{
				LogGroupName:    &logGroupName,
				RetentionInDays: aws.Int32(int32(rid)),
			},
		)

		if err != nil {
			return err
		}
	}

	return nil
}
