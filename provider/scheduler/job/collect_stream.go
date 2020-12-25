package job

import (
	"context"
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
	return "collect_stream"
}

func (s *SequenceAccount) JobTime() string {
	return "0 */5 * * * *"
}

func (s *SequenceAccount) JobMiddleware(job *work.Job, next work.NextMiddlewareFunc) error {
	return next()
}

func (s *SequenceAccount) Retry() uint {
	return 3
}

func (s *SequenceAccount) JobFunc(w *work.Job) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(60)*time.Second)
	defer cancel()

	err := s.streamProvider.CollectAccount(ctx)
	if err != nil {
		return err.Err[0]
	}

	return nil
}
