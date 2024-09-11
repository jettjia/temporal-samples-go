package main

import (
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	"github.com/temporalio/samples-go/a-workflow/approval/workflows"
)

func main() {
	// 创建 Temporal 客户端，每个进程只需要创建一次
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	// 创建一个 worker，指定任务队列名称
	w := worker.New(c, "applicationTaskQueue", worker.Options{})

	// 注册工作流函数和活动函数
	w.RegisterWorkflow(workflows.ApplicationWorkflow)
	w.RegisterActivity(workflows.WaitForApprovalFromB)
	w.RegisterActivity(workflows.WaitForApprovalFromC)

	// 启动 worker，开始监听任务队列
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
