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
	CurrentJob = &Job{
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
	CurrentJob.Status = StatusRunning
	mu.Unlock()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	go runJob(ctx, cancel, CurrentJob)

	SendResponse(w, http.StatusCreated, CurrentJob)
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

	//for _, node := range tree {
	//	if ctx.Err() != nil {
	//		return
	//	}
	//
	//	if isDebug {
	//		color.Yellow("======= Обработка ноды: %s ==========", node.Name)
	//	}
	//
	//	if node.Type == dict.StructTypeFolder {
	//		// какой-то код
	//	} else {
	//		// какой-то код
	//	}
	//
	//	// тут много разного кода
	//}

}

func (h *HeartbeatHandler) GetStatus(w http.ResponseWriter, r *http.Request) {
	SendResponse(w, http.StatusOK, CurrentJob)
}
