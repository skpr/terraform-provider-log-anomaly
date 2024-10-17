package loggroup

import (
	"context"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Create the log group.
func Create(d *schema.ResourceData, m interface{}) error {
	cfg := m.(aws.Config)
	c := cloudwatchlogs.NewFromConfig(cfg)

	var (
		name = d.Get(Name).(string)
	)

	_, err := c.CreateLogGroup(context.TODO(), &cloudwatchlogs.CreateLogGroupInput{
		LogGroupName: aws.String(name),
	})

	if err != nil {
		tflog.Error(context.TODO(), "Error creating log group")
		if !strings.Contains(err.Error(), "ResourceAlreadyExistsException") {
			return err
		}
	}

	lg, _ := findLogGroupByName(context.TODO(), c, name)

	d.Set(Name, name)
	d.SetId(TrimLogGroupARNWildcardSuffix(aws.ToString(lg.Arn)))

	return nil
}
