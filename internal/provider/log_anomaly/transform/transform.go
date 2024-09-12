package transform

import "fmt"

// GetEvaluationFrequency returns the evaluation frequency.
func ToApiValue(evaluationFrequency string) (string, error) {

	switch evaluationFrequency {
	case "5":
		return "FIVE_MIN", nil
	case "10":
		return "TEN_MIN", nil
	case "15":
		return "FIFTEEN_MIN", nil
	case "30":
		return "THIRTY_MIN", nil
	case "60":
		return "ONE_HOUR", nil
	case "":
		return "FIVE_MIN", nil
	default:
		return "", fmt.Errorf("invalid evaluation frequency, use 5, 10, 15, 30 or 60")
	}
}

// transformEvaluationFrequency returns the evaluation frequency.
func FromApiValue(evaluationFrequency string) (string, error) {

	switch evaluationFrequency {
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