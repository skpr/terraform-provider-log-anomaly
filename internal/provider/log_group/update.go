package loggroup

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Update the log group.
func Update(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	cfg := m.(aws.Config)
	c := cloudwatchlogs.NewFromConfig(cfg)

	var (
		name      = d.Get(Name).(string)
		exception *types.ResourceAlreadyExistsException
		// Warning or errors can be collected in a slice type
		diags diag.Diagnostics
	)

	_, err := c.CreateLogGroup(context.TODO(), &cloudwatchlogs.CreateLogGroupInput{
		LogGroupName: aws.String(name),
	})

	if !errors.As(err, &exception) && err != nil {
		return diag.FromErr(err)
	}

	lg, err := findLogGroupByName(ctx, c, name)
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set(Name, name)
	d.SetId(TrimLogGroupARNWildcardSuffix(aws.ToString(lg.Arn)))

	return diags
}
