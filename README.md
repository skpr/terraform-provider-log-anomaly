Terraform Provider: Log Anomaly Detector
=============================

Terraform provider for deploying a log anomaly detector on a log group

## Usage

```hcl
resource "aws_cloudwatch_log_anomaly_detector" "prod_logs" {
    name = "prod-logs"
    log_group = "arn:aws:logs:ap-southeast-2:022593270620:log-group:nmhp/vpc/flow"
    evaluation_frequency = "15"
}
```
