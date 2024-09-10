package activities

import (
	"context"

	"go.temporal.io/sdk/activity"
)

// GetUser is the implementation for Temporal activity
func GetUser(ctx context.Context) (string, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("GetUser activity called")
	return "Temporal", nil
}
