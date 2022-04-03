package main

import (
	"log"
	"time"

	"github.com/hibiken/asynq"
	"github.com/iradukunda1/asynq-test/task"
)

// client.go
func main() {
	client := asynq.NewClient(asynq.RedisClientOpt{Addr: "localhost:6379"})

	t1, err := task.NewWelcomeEmailTask(42)
	if err != nil {
		log.Fatal(err)
	}

	t2, err := task.NewReminderEmailTask(42)
	if err != nil {
		log.Fatal(err)
	}

	// Process the task immediately.
	info, err := client.Enqueue(t1)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf(" [*] Successfully enqueued task: %+v", info)

	// Process the task 24 hours later.
	info, err = client.Enqueue(t2, asynq.ProcessIn(2*time.Second))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf(" [*] Successfully enqueued task: %+v", info)
}
