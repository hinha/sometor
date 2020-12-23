package job

import (
	"context"
	"fmt"
	"github.com/gocraft/work"
	"github.com/hinha/sometor/provider"
	"time"
)

type SequenceAccount struct {
	streamProvider provider.TwitterStreaming
}

func NewSequenceAccount(provider provider.TwitterStreaming) *SequenceAccount {
	return &SequenceAccount{streamProvider: provider}
}

func (s *SequenceAccount) JobName() string {
	return "collect_account"
}

func (s *SequenceAccount) JobTime() string {
	return "@every 0h2m0s"
}

func (s *SequenceAccount) JobMiddleware(job *work.Job, next work.NextMiddlewareFunc) error {
	fmt.Println("Starting job: ", job.Name)
	return next()
}

func (s *SequenceAccount) Retry() uint {
	return 3
}

func (s *SequenceAccount) JobFunc(w *work.Job) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(60)*time.Second)
	defer cancel()

	fmt.Println(s.streamProvider.CollectAccount(ctx))

	return nil
}
