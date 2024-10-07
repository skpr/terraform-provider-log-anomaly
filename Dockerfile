FROM hashicorp/terraform:1.2.5 as run

RUN apk add bash

RUN mkdir -p /root/.terraform.d/plugins

ADD ./bin/terraform-provider-loganomaly_linux-amd64 /root/.terraform.d/plugins/terraform.local/skpr/log-anomaly-detector/99.0.0/linux_amd64/terraform-provider-log-anomaly-detector_linux-amd64_v99.0.0

RUN chmod +x /root/.terraform.d/plugins/terraform.local/*/*/*/linux_amd64/terraform-provider-*
