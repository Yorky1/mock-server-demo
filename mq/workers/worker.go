package workers

import (
	"fmt"
	"time"
)

type Worker struct {
	task_ch <-chan string
}

func (w *Worker) Routine() {
	for task := range w.task_ch {
		LongOperation(task)
	}
}

func LongOperation(task string) {
	// long operation: db/logs/queries
	fmt.Printf("get task: %s, starting... ", task)
	time.Sleep(5 * time.Second)
	fmt.Println("done")
}

func CreateWorker(ch <-chan string) Worker {
	return Worker{
		task_ch: ch,
	}
}
