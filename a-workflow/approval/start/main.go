package main

import (
	"context"
	"log"

	"go.temporal.io/sdk/client"

	"github.com/temporalio/samples-go/a-workflow/approval/workflows"
)

func main() {
	// 创建Temporal客户端
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	// 定义工作流选项
	workflowOptions := client.StartWorkflowOptions{
		ID:        "application-workflow-id",
		TaskQueue: "applicationTaskQueue",
	}

	// 启动工作流执行
	we, err := c.ExecuteWorkflow(context.Background(), workflowOptions, workflows.ApplicationWorkflow)
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
	}
	log.Println("Started workflow", "WorkflowID", we.GetID(), "RunID", we.GetRunID())

	// 等待工作流完成
	var result workflows.Application
	err = we.Get(context.Background(), &result)
	if err != nil {
		log.Fatalln("Unable to get workflow result", err)
	}

	// 输出结果
	log.Println("Workflow result:", result.Status)
}
