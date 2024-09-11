package main

import (
	"log"

	"github.com/temporalio/samples-go/a-workflow/dynamic_approve/dynamic"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	// 创建 Temporal 客户端
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	// 创建 Temporal 工作者
	w := worker.New(c, "reviewTaskQueue", worker.Options{})

	// 注册工作流
	w.RegisterWorkflow(dynamic.DynamicReviewWorkflow)

	// 注册活动
	w.RegisterActivity(dynamic.ApproveRequest)
	w.RegisterActivity(dynamic.RejectRequest)

	// 启动工作者
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to run worker", err)
	}
}
