package command

import (
	"context"
	"fmt"
	"github.com/hinha/sometor/provider"
	"time"
)

// Run is a command to run api engine
type Run struct {
	engine provider.ScheduleEngine
}

// NewRunScheduler return CLI to run scheduler
func NewRunScheduler(engine provider.ScheduleEngine) *Run {
	return &Run{engine: engine}
}

// Use return how the command used
func (r *Run) Use() string {
	return "run:cron"
}

// Example of the command
func (r *Run) Example() string {
	return "run:cron"
}

// Short description about the command
func (r *Run) Short() string {
	return "Run Scheduler cron engine"
}

func (r *Run) Run(args []string) {
	r.engine.Run()

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	r.engine.Shutdown(ctx)
	fmt.Println("\nGracefully shutdown the scheduler...")
}
