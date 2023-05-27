package workers

func CreateWorkerPool(worker_cnt uint64) chan<- string {
	task_ch := make(chan string, worker_cnt)

	for i := 0; i < int(worker_cnt); i += 1 {
		worker := CreateWorker(task_ch)
		go func() {
			worker.Routine()
		}()
	}

	return task_ch
}
