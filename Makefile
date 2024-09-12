#!/usr/bin/make -f

export CGO_ENABLED=0

define build_step
	GOOS=$(1) GOARCH=$(2) go build -o bin/terraform-provider-loganomaly_$(1)-$(2) -ldflags='-extldflags "-static"' github.com/skpr/terraform-provider-log-anomaly/cmd/terraform-provider-log-anomaly
endef

# Builds the project.
build:
	$(call build_step,linux,amd64)
	$(call build_step,linux,arm64)

# Run all lint checking with exit codes for CI.
lint:
	golint -set_exit_status `go list ./... | grep -v /vendor/`

# Run tests with coverage reporting.
test:
	go test -cover ./...

.PHONY: *