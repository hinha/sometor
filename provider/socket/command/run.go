package command

import (
	"context"
	"fmt"
	"github.com/hinha/sometor/provider"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Run is a command to run api engine
type Run struct {
	engine provider.SocketEngine
}

// NewRun return CLI to run api engine
func NewRunSocket(engine provider.SocketEngine) *Run {
	return &Run{engine: engine}
}

// Use return how the command used
func (r *Run) Use() string {
	return "run:socket"
}

// Example of the command
func (r *Run) Example() string {
	return "run:socket"
}

// Short description about the command
func (r *Run) Short() string {
	return "Run Socket engine"
}

func (r *Run) Run(args []string) {
	go func() {
		_ = r.engine.Run()
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 3 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// omit the error
	_ = r.engine.Shutdown(ctx)

	fmt.Println("\nGracefully shutdown the server...")
}
