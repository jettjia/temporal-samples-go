package main

import (
	"go.uber.org/zap"

	"github.com/temporalio/samples-go/a-workflow/activities/activities"
	"github.com/temporalio/samples-go/a-workflow/activities/workflows"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	logger.Info("Zap logger created")

	// The client is a heavyweight object that should be created once
	serviceClient, err := client.NewClient(client.Options{
		HostPort:  client.DefaultHostPort,
		Namespace: client.DefaultNamespace,
	})
	if err != nil {
		logger.Fatal("Unable to start worker", zap.Error(err))
	}

	worker := worker.New(serviceClient, "tutorial_tq", worker.Options{})

	worker.RegisterWorkflow(workflows.Greetings)
	worker.RegisterActivity(activities.GetUser)
	worker.RegisterActivity(activities.SendGreeting)

	err = worker.Start()
	if err != nil {
		logger.Fatal("Unable to start worker", zap.Error(err))
	}

	select {}
}
