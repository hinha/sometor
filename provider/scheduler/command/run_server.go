package command

import (
	"context"
	"fmt"
	"github.com/hinha/sometor/provider"
	"time"
)

// Run is a command to run api engine
type RunSchedulerServer struct {
	engine provider.ScheduleEngine
}

// NewRunSchedulerServer return CLI to run scheduler
func NewRunSchedulerServer(engine provider.ScheduleEngine) *RunSchedulerServer {
	return &RunSchedulerServer{engine: engine}
}

// Use return how the command used
func (r *RunSchedulerServer) Use() string {
	return "run:cron:server"
}

// Example of the command
func (r *RunSchedulerServer) Example() string {
	return "run:cron:server"
}

// Short description about the command
func (r *RunSchedulerServer) Short() string {
	return "Run Scheduler cron engine"
}

func (r *RunSchedulerServer) Run(args []string) {
	r.engine.Run()

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	r.engine.Shutdown(ctx)
	fmt.Println("\nGracefully shutdown the scheduler...")
}
