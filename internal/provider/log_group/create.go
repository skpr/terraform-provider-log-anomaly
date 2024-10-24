package loggroup

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Create the log group.
func Create(d *schema.ResourceData, m interface{}) error {
	cfg := m.(aws.Config)
	c := cloudwatchlogs.NewFromConfig(cfg)

	var (
		name      = d.Get(Name).(string)
		ctx       = context.TODO()
		exception *types.ResourceAlreadyExistsException
	)

	_, err := c.CreateLogGroup(context.TODO(), &cloudwatchlogs.CreateLogGroupInput{
		LogGroupName: aws.String(name),
	})

	if !errors.As(err, &exception) && err != nil {
		return err
	}

	lg, err := findLogGroupByName(ctx, c, name)
	if err != nil {
		return err
	}

	d.Set(Name, name)
	d.SetId(TrimLogGroupARNWildcardSuffix(aws.ToString(lg.Arn)))

	return nil
}
