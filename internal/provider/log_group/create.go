package loggroup

import (
	"context"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Create the log anomaly detector.
func Create(d *schema.ResourceData, m interface{}) error {
	cfg := m.(aws.Config)
	c := cloudwatchlogs.NewFromConfig(cfg)

	var (
		name            = d.Get(Name).(string)
		retentionInDays = d.Get(RetentionInDays).(int)
		// region          = d.Get(Region).(string)
		// accountID       = d.Get(AccountID).(string)
	)

	if retentionInDays < 1 {
		retentionInDays = 90
	}

	_, err := c.CreateLogGroup(context.TODO(), &cloudwatchlogs.CreateLogGroupInput{
		LogGroupName: aws.String(name),
	})

	if err != nil {
		tflog.Error(context.TODO(), "Error creating log group")
		if !strings.Contains(err.Error(), "ResourceAlreadyExistsException") {
			return err
		}
	}

	_, err = c.PutRetentionPolicy(context.TODO(), &cloudwatchlogs.PutRetentionPolicyInput{
		LogGroupName:    aws.String(name),
		RetentionInDays: aws.Int32(int32(retentionInDays)),
	})

	if err != nil {
		return err
	}

	lg, _ := findLogGroupByName(context.TODO(), c, name)

	d.Set("log_group_name", name)
	d.SetId(TrimLogGroupARNWildcardSuffix(aws.ToString(lg.Arn)))

	return nil
}
