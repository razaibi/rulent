package logic

import (
	"fmt"
	"reflect"
	"rulent/models"
	"strconv"
	"strings"
)

func ValidateJSON(payload map[string]interface{}, config *models.Config) ([]models.Outcome, bool) {
	eventNames, ok := payload["events"].([]interface{})
	if !ok {
		return nil, false // 'events' not present or not an array
	}

	for _, eventName := range eventNames {
		eventNameStr, ok := eventName.(string)
		if !ok {
			continue // Skip this entry if it's not a string
		}

		event, ok := config.Events[eventNameStr]
		if !ok {
			continue // Skip if the event name doesn't exist in the configuration
		}

		if validateEvent(payload, event) {
			return event.Outcomes, true // Return outcomes if the event rules are satisfied
		}
	}

	return nil, false // No matching events or none of the events' rules were satisfied
}

func validateEvent(payload map[string]interface{}, event models.Event) bool {
	allRulesSatisfied := true

	for _, rule := range event.Rules {
		ruleSatisfied := validateRule(payload, rule.Conditions)

		if event.RulesOperator == "or" && ruleSatisfied {
			return true // If 'or' operator is used, return true if any rule is satisfied
		} else if event.RulesOperator == "and" && !ruleSatisfied {
			allRulesSatisfied = false // If 'and' operator, all rules must be satisfied
		}
	}

	if event.RulesOperator == "and" {
		return allRulesSatisfied // For 'and', return true only if all rules are satisfied
	}
	return false // Default return false for 'or' if none of the rules are satisfied
}

func validateRule(payload map[string]interface{}, conditions []map[string]interface{}) bool {
	for _, condition := range conditions {
		for key, value := range condition {
			keys := strings.Split(key, ".")
			if !validateField(payload, keys, value) {
				return false
			}
		}
	}
	return true
}

func validateField(current map[string]interface{}, keys []string, expectedValue interface{}) bool {
	if len(keys) == 0 {
		return false
	}

	key := keys[0]
	if len(keys) == 1 {
		actualValue, ok := current[key]
		if !ok {
			return false // Key does not exist
		}

		switch v := expectedValue.(type) {
		case string:
			// Check for conditions starting with > or <
			if strings.HasPrefix(v, ">") ||
				strings.HasPrefix(v, "<") ||
				strings.HasPrefix(v, "!") {
				return compareValues(v, actualValue)
			}
		}

		// Fallback to direct equality if not a special condition
		return reflect.DeepEqual(actualValue, expectedValue)
	}

	next, ok := current[key].(map[string]interface{})
	if !ok {
		return false
	}
	return validateField(next, keys[1:], expectedValue)
}

func compareValues(condition string, actualValue interface{}) bool {
	operator := condition[:1]
	expectedValStr := condition[1:]

	// Convert the expected value string to float64 for comparison
	expectedVal, err := strconv.ParseFloat(expectedValStr, 64)
	if err != nil {
		fmt.Printf("Error parsing expected value '%s' to float64: %v\n", expectedValStr, err)
		return false
	}

	var actualValFloat float64
	switch actual := actualValue.(type) {
	case float64:
		actualValFloat = actual
	case float32:
		actualValFloat = float64(actual)
	case int:
		actualValFloat = float64(actual)
	case int32:
		actualValFloat = float64(actual)
	case int64:
		actualValFloat = float64(actual)
	case string:
		// Try to convert a string to float64
		var convErr error
		actualValFloat, convErr = strconv.ParseFloat(actual, 64)
		if convErr != nil {
			fmt.Printf("Error converting string '%s' to float64: %v\n", actual, convErr)
			return false
		}
	default:
		fmt.Printf("Actual value '%v' is of type %T, which is not supported for comparison\n", actualValue, actualValue)
		return false
	}

	// Perform the comparison based on the operator
	switch operator {
	case ">":
		return actualValFloat > expectedVal
	case "<":
		return actualValFloat < expectedVal
	case "!":
		return actualValFloat != expectedVal
	default:
		fmt.Println("Unsupported operator:", operator)
		return false
	}
}
