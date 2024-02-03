package logic

import (
	"rulent/models"
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
		return current[key] == expectedValue
	}

	next, ok := current[key].(map[string]interface{})
	if !ok {
		return false
	}
	return validateField(next, keys[1:], expectedValue)
}
