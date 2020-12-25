package scheduler

import (
	"context"
	"fmt"
	"github.com/gocraft/work"
	"github.com/gomodule/redigo/redis"
	"github.com/hinha/sometor/provider"
	"github.com/hinha/sometor/provider/scheduler/command"
	"os"
	"os/signal"
)

type Server struct {
	namespace string
	work      *work.WorkerPool
}

func FabricateServer(namespace string) *Server {
	redisPool := &redis.Pool{
		MaxActive: 5,
		MaxIdle:   5,
		Wait:      true,
		Dial:      func() (redis.Conn, error) { return redis.Dial("tcp", fmt.Sprintf("%s:6379", os.Getenv("REDIS_HOST"))) },
	}

	return &Server{namespace: namespace, work: work.NewWorkerPool(struct{}{}, 10, namespace, redisPool)}
}

// FabricateCommand insert schedule related command
func (s *Server) FabricateCommand(cmd provider.Command) {
	cmd.InjectCommand(
		command.NewRunSchedulerServer(s),
	)
}

func (s *Server) Run() {
	// Start processing jobs
	s.work.Start()
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, os.Kill)
	<-signalChan
}

func (s *Server) Inject(handler provider.ScheduleHandler) {
	s.work.PeriodicallyEnqueue(handler.JobTime(), handler.JobName())
	s.work.Middleware(handler.JobMiddleware)
	s.work.JobWithOptions(handler.JobName(), work.JobOptions{Priority: 10, MaxFails: handler.Retry()}, handler.JobFunc)
}

func (s *Server) Shutdown(ctx context.Context) {
	s.work.Stop()
}
