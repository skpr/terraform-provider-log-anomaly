package loggroup

import (
	// "github.com/aws/aws-sdk-go-v2/aws/session"

	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Delete the log group
func Delete(d *schema.ResourceData, m interface{}) error {
	cfg := m.(aws.Config)
	c := cloudwatchlogs.NewFromConfig(cfg)

	LogGroupName := d.Get(Name).(string)

	_, err := c.DeleteLogGroup(
		context.TODO(),
		&cloudwatchlogs.DeleteLogGroupInput{
			LogGroupName: &LogGroupName,
		},
	)
	if err != nil {
		return err
	}

	return nil
}
