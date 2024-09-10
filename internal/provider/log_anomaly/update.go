package loganomaly

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/hashicorp/terraform/helper/schema"
)

// Update the Project.
func Update(d *schema.ResourceData, m interface{}) error {
	sess := m.(*session.Session)

	var (
		name                = d.Get(Name).(string)
		logGroup            = d.Get(LogGroup).(string)
		evaluationFrequency = d.Get(EvaluationFrequency).(string)
	)

	c := cloudwatchlogs.New(sess)

	_, err := c.CreateLogAnomalyDetector(&cloudwatchlogs.CreateLogAnomalyDetectorInput{
		DetectorName:        aws.String(name),
		LogGroupArnList:     aws.StringSlice([]string{logGroup}),
		EvaluationFrequency: aws.String(evaluationFrequency),
	})

	if err != nil {
		return err
	}

	return nil

	// source, err := http.Get(url)
	// if err != nil {
	// 	return err
	// }

	// if source.StatusCode != http.StatusOK {
	// 	return fmt.Errorf("failed to get source: %s", source.Status)
	// }

	// bodyBytes, err := io.ReadAll(source.Body)
	// if err != nil {
	// 	return fmt.Errorf("failed to read source body: %w", err)
	// }

	// err = source.Body.Close()
	// if err != nil {
	// 	return fmt.Errorf("failed to close source body: %w", err)
	// }

	// bodyString := string(bodyBytes)

	// hashHave := md5.GetHash(bodyString)

	// if hashHave != hashWant {
	// 	return fmt.Errorf("hash mismatch: %s != %s", hashHave, hashWant)
	// }

	// _, err = s3manager.NewUploader(sess).Upload(&s3manager.UploadInput{
	// 	Bucket: aws.String(bucket),
	// 	Key:    aws.String(key),
	// 	Body:   bytes.NewReader(bodyBytes),
	// })
	// if err != nil {
	// 	return err
	// }

	// d.SetId(id.Join(bucket, key))

	// return nil
}
