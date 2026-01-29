package handler

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

const (
	StatusIdle    = "idle"
	StatusRunning = "running"
)

type Job struct {
	Status string `json:"status"`
}

var (
	currentJob = &Job{
		Status: StatusIdle,
	}
	mu sync.Mutex
)

type JobHandler struct {
}

func NewJobHandler() *HeartbeatHandler {
	return &HeartbeatHandler{}
}

func (h *HeartbeatHandler) Start(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	currentJob.Status = StatusRunning
	mu.Unlock()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	go runJob(ctx, cancel, currentJob)

	SendResponse(w, http.StatusCreated, currentJob)
}

func (h *HeartbeatHandler) GetStatus(w http.ResponseWriter, r *http.Request) {
	SendResponse(w, http.StatusOK, currentJob)
}

func runJob(ctx context.Context, cancel context.CancelFunc, job *Job) {
	defer cancel()
	defer func() {
		mu.Lock()
		job.Status = StatusIdle
		fmt.Println("stop job")
		mu.Unlock()
	}()

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	job.Status = StatusRunning // имитация обновления
	fmt.Println("tick job")

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			mu.Lock()
			job.Status = StatusRunning // имитация обновления
			fmt.Println("tick job")
			mu.Unlock()
		}
	}
}
