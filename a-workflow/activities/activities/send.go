package activities

import (
	"context"
	"fmt"

	"go.temporal.io/sdk/activity"
)

// SendGreeting is the implementation for Temporal activity
func SendGreeting(ctx context.Context, user string) error {
	logger := activity.GetLogger(ctx)
	logger.Info("SendGreeting activity called")

	fmt.Printf("Greeting sent to user: %v\n", user)
	return nil
}
