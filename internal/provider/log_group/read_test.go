package loggroup

import (
	"testing"
)

func TestTrimLogGroupARNWildcardSuffix(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		arn  string
		want string
	}{
		{
			name: "ARN with wildcard suffix",
			arn:  "arn:aws:logs:us-west-2:123456789012:log-group:my-log-group:*",
			want: "arn:aws:logs:us-west-2:123456789012:log-group:my-log-group",
		},
		{
			name: "ARN without wildcard suffix",
			arn:  "arn:aws:logs:us-west-2:123456789012:log-group:my-log-group",
			want: "arn:aws:logs:us-west-2:123456789012:log-group:my-log-group",
		},
		{
			name: "Empty ARN",
			arn:  "",
			want: "",
		},
		{
			name: "ARN with multiple wildcard suffixes",
			arn:  "arn:aws:logs:us-west-2:123456789012:log-group:my-log-group:*:*",
			want: "arn:aws:logs:us-west-2:123456789012:log-group:my-log-group:*",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TrimLogGroupARNWildcardSuffix(tt.arn); got != tt.want {
				t.Errorf("TrimLogGroupARNWildcardSuffix() = %v, want %v", got, tt.want)
			}
		})
	}
}
