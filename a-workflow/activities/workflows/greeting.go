package workflows

import (
	"time"

	"github.com/temporalio/samples-go/a-workflow/activities/activities"
	"go.temporal.io/sdk/workflow"
	"go.uber.org/zap"
)

// Greetings is the implementation for Temporal workflow
func Greetings(ctx workflow.Context) error {
	logger := workflow.GetLogger(ctx)
	logger.Info("Workflow Greetings started")

	ao := workflow.ActivityOptions{
		ScheduleToStartTimeout: time.Hour,
		StartToCloseTimeout:    time.Hour,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	var user string
	err := workflow.ExecuteActivity(ctx, activities.GetUser).Get(ctx, &user)
	if err != nil {
		return err
	}

	err = workflow.ExecuteActivity(ctx, activities.SendGreeting, user).Get(ctx, nil)
	if err != nil {
		return err
	}

	logger.Info("Greetings workflow complete", zap.String("user", user))
	return nil
}
