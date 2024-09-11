package main

import (
	"context"
	"log"
	"math/rand"
	"time"

	"github.com/temporalio/samples-go/a-workflow/dynamic_approve/dynamic"

	"go.temporal.io/sdk/client"
)

func main() {
	// 创建 Temporal 客户端
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	// 定义工作流选项
	workflowOptions := client.StartWorkflowOptions{
		ID:        "review-workflow-id",
		TaskQueue: "reviewTaskQueue",
	}

	// 生成随机数
	rand.Seed(time.Now().UnixNano()) // 设置随机数种子，以当前时间作为种子保证随机性
	arg := rand.Intn(100) + 1        // 生成 1 到 100 的随机数
	// 启动工作流执行
	we, err := c.ExecuteWorkflow(context.Background(), workflowOptions, dynamic.DynamicReviewWorkflow, arg)
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
	}
	log.Println("Started workflow", "WorkflowID", we.GetID(), "RunID", we.GetRunID())

	// 等待工作流完成
	var result string
	err = we.Get(context.Background(), &result)
	if err != nil {
		log.Fatalln("Unable to get workflow result", err)
	}

	// 输出结果
	log.Println("Workflow result:", result)
}
