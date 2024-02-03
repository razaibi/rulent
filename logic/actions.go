package logic

import (
	"fmt"
	"rulent/models"
	"sync"
)

type ActionFunc func(payload map[string]interface{}, parameters map[string]interface{}) error

var actionFuncs = map[string]ActionFunc{
	"email": emailAction,
	"log":   logAction,
}

func emailAction(payload map[string]interface{}, parameters map[string]interface{}) error {
	fmt.Println("Email action.")
	return nil
}

func logAction(payload map[string]interface{}, parameters map[string]interface{}) error {
	fmt.Println("Logging action.")
	return nil
}

func ExecuteActions(outcomes []models.Outcome, payload map[string]interface{}, errorChan chan<- error, wg *sync.WaitGroup) {
	for _, outcome := range outcomes {
		actionFunc, ok := actionFuncs[outcome.Action]
		if !ok {
			errorChan <- fmt.Errorf("action '%s' is not supported", outcome.Action)
			continue
		}

		// Execute the action synchronously or asynchronously based on the outcome mode
		if outcome.Mode == "async" {
			wg.Add(1) // Increment the wait group counter for async actions
			go func(outcome models.Outcome) {
				defer wg.Done() // Decrement the counter when the goroutine completes
				err := actionFunc(payload, outcome.Parameters)
				if err != nil {
					errorChan <- fmt.Errorf("error executing async action '%s': %v", outcome.Action, err)
				}
			}(outcome)
		} else { // synchronous execution
			err := actionFunc(payload, outcome.Parameters)
			if err != nil {
				errorChan <- fmt.Errorf("error executing sync action '%s': %v", outcome.Action, err)
			}
		}
	}
}
