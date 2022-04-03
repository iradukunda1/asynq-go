package main

import (
	"log"

	"github.com/hibiken/asynq"
	"github.com/iradukunda1/asynq-test/task"
)

// Task payload for any email related tasks.
type EmailTaskPayload struct {
	// ID for the email recipient.
	UserID int
}

func main() {
	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: "localhost:6379"},
		asynq.Config{Concurrency: 10},
	)

	mux := asynq.NewServeMux()
	mux.HandleFunc(task.TypeWelcomeEmail, task.HandleWelcomeEmailTask)
	mux.HandleFunc(task.TypeReminderEmail, task.HandleReminderEmailTask)

	if err := srv.Run(mux); err != nil {
		log.Fatal(err)
	}
}
