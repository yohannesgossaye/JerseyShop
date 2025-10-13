package main

import (
	"log"

	workflows "github.com/yohannesgossaye/pkgs/temporalworkflow"

	"github.com/joho/godotenv"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf(" Failed to load .env: %v", err)
	} else {
		log.Println("✅ .env file loaded successfully")
	}

	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create Temporal client", err)
	}
	defer c.Close()

	w := worker.New(c, "USER_TASK_QUEUE_V2", worker.Options{})

	w.RegisterWorkflow(workflows.UserSignupWorkflow)
	w.RegisterActivity(workflows.SaveUserActivity)
	w.RegisterActivity(workflows.SendEmailActivity)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
