package loggroup

import (
	"context"
	"log"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	// "github.com/aws/aws-sdk-go-v2/aws/session"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Read the log anomaly detector.
func Read(d *schema.ResourceData, m interface{}) error {
	cfg := m.(aws.Config)
	c := cloudwatchlogs.NewFromConfig(cfg)

	logGroupName := d.Get("name").(string)

	lg, err := findLogGroupByName(context.TODO(), c, logGroupName)

	if !d.IsNewResource() && err != nil {
		log.Printf("[WARN] CloudWatch Logs Log Group (%s) not found, removing from state", d.Id())
		d.SetId("")
		return err
	}
	// logGroupName = TrimLogGroupARNWildcardSuffix(aws.ToString(lg.Arn))

	d.SetId(TrimLogGroupARNWildcardSuffix(aws.ToString(lg.Arn)))
	d.Set(Name, logGroupName)
	d.Set(RetentionInDays, lg.RetentionInDays)

	return nil
}

func findLogGroupByName(ctx context.Context, conn *cloudwatchlogs.Client, name string) (*types.LogGroup, error) {
	input := &cloudwatchlogs.DescribeLogGroupsInput{
		LogGroupNamePrefix: aws.String(name),
	}

	pages := cloudwatchlogs.NewDescribeLogGroupsPaginator(conn, input)
	for pages.HasMorePages() {
		page, err := pages.NextPage(ctx)

		if err != nil {
			return nil, err
		}

		for _, v := range page.LogGroups {
			if aws.ToString(v.LogGroupName) == name {
				return &v, nil
			}
		}
	}

	return nil, nil
}

// TrimLogGroupARNWildcardSuffix removes the wildcard suffix from a log group ARN.
func TrimLogGroupARNWildcardSuffix(arn string) string {
	logGroupARNWildcardSuffix := ":*"
	return strings.TrimSuffix(arn, logGroupARNWildcardSuffix)
}
