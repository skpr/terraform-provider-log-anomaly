package transform

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
)

// ToAPIValue returns the evaluation frequency.
func ToAPIValue(evaluationFrequency string) (types.EvaluationFrequency, error) {

	switch evaluationFrequency {
	case "1":
		return types.EvaluationFrequencyOneMin, nil
	case "5":
		return types.EvaluationFrequencyFiveMin, nil
	case "10":
		return types.EvaluationFrequencyTenMin, nil
	case "15":
		return types.EvaluationFrequencyFifteenMin, nil
	case "30":
		return types.EvaluationFrequencyThirtyMin, nil
	case "60":
		return types.EvaluationFrequencyOneHour, nil
	case "":
		return types.EvaluationFrequencyFiveMin, nil
	default:
		return "", fmt.Errorf("invalid evaluation frequency, use 5, 10, 15, 30 or 60")
	}
}

// FromAPIValue returns the evaluation frequency.
func FromAPIValue(evaluationFrequency types.EvaluationFrequency) (string, error) {

	switch evaluationFrequency {
	case "ONE_MIN":
		return "1", nil
	case "FIVE_MIN":
		return "5", nil
	case "TEN_MIN":
		return "10", nil
	case "FIFTEEN_MIN":
		return "15", nil
	case "THIRTY_MIN":
		return "30", nil
	case "ONE_HOUR":
		return "60", nil
	default:
		return "", fmt.Errorf("invalid evaluation frequency, use FIVE_MIN, TEN_MIN, FIFTEEN_MIN, THIRTY_MIN or ONE_HOUR")
	}
}
