package dynamic

import (
	"time"

	"go.temporal.io/sdk/workflow"
)

// DynamicReviewWorkflow 是一个动态审核流程的工作流
func DynamicReviewWorkflow(ctx workflow.Context, requestID int) (string, error) {

	// 定义活动选项
	activityOptions := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 10,
	}
	ctx = workflow.WithActivityOptions(ctx, activityOptions)

	// 动态决定审核流程的步骤
	if requestID%2 == 0 {
		// 如果请求 ID 是偶数，执行 ApproveRequest 活动
		err := workflow.ExecuteActivity(ctx, ApproveRequest, requestID).Get(ctx, new(bool))
		if err != nil {
			return "", err
		}

		return "Request approved", nil

	} else {
		// 如果请求 ID 是奇数，执行 RejectRequest 活动
		err := workflow.ExecuteActivity(ctx, RejectRequest, requestID).Get(ctx, new(bool))
		if err != nil {
			return "", err
		}

		return "Request rejected", nil
	}

	return "Unknown state", nil
}
