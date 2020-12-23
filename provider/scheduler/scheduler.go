package scheduler

import (
	"context"
	"fmt"
	"github.com/gocraft/work"
	"github.com/gomodule/redigo/redis"
	"github.com/hinha/sometor/provider"
	"github.com/hinha/sometor/provider/scheduler/command"
	"github.com/hinha/sometor/provider/scheduler/job"
	"os"
	"os/signal"
)

type Scheduler struct {
	namespace string
	work      *work.WorkerPool
}

func Fabricate(namespace string) *Scheduler {
	redisPool := &redis.Pool{
		MaxActive: 5,
		MaxIdle:   5,
		Wait:      true,
		Dial:      func() (redis.Conn, error) { return redis.Dial("tcp", fmt.Sprintf("%s:6379", os.Getenv("REDIS_HOST"))) },
	}

	return &Scheduler{namespace: namespace, work: work.NewWorkerPool(struct{}{}, 10, namespace, redisPool)}
}

// FabricateCommand insert schedule related command
func (s *Scheduler) FabricateCommand(cmd provider.Command) {
	cmd.InjectCommand(
		command.NewRunScheduler(s),
	)
}

func (s *Scheduler) Run() {

	s.Inject(job.NewPingDB())

	// Start processing jobs
	s.work.Start()
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, os.Kill)
	<-signalChan
}

func (s *Scheduler) Inject(handler provider.ScheduleHandler) {
	s.work.PeriodicallyEnqueue(handler.JobTime(), handler.JobName())
	s.work.Middleware(handler.JobMiddleware)
	s.work.JobWithOptions(handler.JobName(), work.JobOptions{Priority: 10, MaxFails: handler.Retry()}, handler.JobFunc)
}

func (s *Scheduler) Shutdown(ctx context.Context) {
	s.work.Stop()
}
