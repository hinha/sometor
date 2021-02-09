package job

import (
	"context"
	"github.com/gocraft/work"
	"github.com/hinha/sometor/provider"
	"time"
)

type SequenceAccount struct {
	streamProvider provider.AllProviderStreaming
}

func NewSequenceAccount(provider provider.AllProviderStreaming) *SequenceAccount {
	return &SequenceAccount{streamProvider: provider}
}

func (s *SequenceAccount) JobName() string {
	return "collect_stream_scraping"
}

func (s *SequenceAccount) JobTime() string {
	return "0 */10 * * * *" // Scraping Local
}

func (s *SequenceAccount) JobMiddleware(job *work.Job, next work.NextMiddlewareFunc) error {
	return next()
}

func (s *SequenceAccount) Retry() uint {
	return 3
}

// handle scraping to upload object
func (s *SequenceAccount) JobFunc(w *work.Job) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(2)*time.Minute)
	defer cancel()

	err := s.streamProvider.CollectAccount(ctx)
	if err != nil {
		return err.Err[0]
	}

	return nil
}
