package workflows

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.temporal.io/sdk/workflow"
)

const (
	applicationWorkflowName = "applicationWorkflow"
)

// Application 结构体表示申请对象，包含申请的状态
type Application struct {
	Status string
}

// Approval 结构体表示审批结果，包含是否批准的布尔值
type Approval struct {
	Approved bool
}

// ApplicationWorkflow 是主要的工作流函数，处理申请流程
func ApplicationWorkflow(ctx workflow.Context) (*Application, error) {
	// 设置活动的超时时间选项
	ao := workflow.ActivityOptions{
		ScheduleToStartTimeout: time.Minute,
		StartToCloseTimeout:    time.Minute,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	// 创建一个初始状态为“Pending Approval by B and C”的申请对象
	var app Application
	app.Status = "Pending Approval by B and C"

	log.Println("Starting application workflow")

	// 执行一个活动，模拟申请的发起，这里只是打印一条日志
	err := workflow.ExecuteActivity(ctx, func(ctx context.Context) error {
		log.Println("Application initiated.")
		return nil
	}).Get(ctx, nil)
	if err != nil {
		return nil, err
	}

	log.Println("Application initiation completed.")

	// 等待 B 的审批结果，结果存储在 bApproval 变量中
	bApproval := false
	err = workflow.ExecuteActivity(ctx, WaitForApprovalFromB).Get(ctx, &bApproval)
	if err != nil {
		return nil, err
	}

	log.Println("B approval result:", bApproval)

	// 如果 B 批准了申请
	if bApproval {
		// 等待 C 的审批结果，结果存储在 cApproval 变量中
		cApproval := false
		err = workflow.ExecuteActivity(ctx, WaitForApprovalFromC).Get(ctx, &cApproval)
		if err != nil {
			return nil, err
		}

		log.Println("C approval result:", cApproval)

		// 如果 C 也批准了申请
		if cApproval {
			app.Status = "Approved"
		} else {
			// 如果 C 驳回了申请
			app.Status = "Rejected by C"
		}
	} else {
		// 如果 B 驳回了申请
		app.Status = "Rejected by B"
	}

	log.Println("Application workflow completed with status:", app.Status)
	return &app, nil
}

// WaitForApprovalFromB 模拟等待 B 的审批，随机返回批准或驳回
func WaitForApprovalFromB(ctx context.Context) error {
	// 在实际场景中，这可能是一个外部调用或用户交互。这里只是为了简单起见随机模拟审批结果
	if time.Now().UnixNano()%2 == 0 {
		log.Println("Approved by B")
		return nil
	} else {
		log.Println("Rejected by B")
		return fmt.Errorf("rejected by B")
	}
}

// WaitForApprovalFromC 模拟等待 C 的审批，随机返回批准或驳回
func WaitForApprovalFromC(ctx context.Context) error {
	// 与 WaitForApprovalFromB 类似，随机模拟审批结果
	if time.Now().UnixNano()%2 == 0 {
		log.Println("Approved by C")
		return nil
	} else {
		log.Println("Rejected by C")
		return fmt.Errorf("rejected by C")
	}
}
