package job

import (
	"context"
	"github.com/gocraft/work"
	"github.com/hinha/sometor/provider"
	"time"
)

type CollectStreamObject struct {
	streamProvider provider.AllProviderStreaming
}

func NewCollectStream(provider provider.AllProviderStreaming) *CollectStreamObject {
	return &CollectStreamObject{streamProvider: provider}
}

func (s *CollectStreamObject) JobName() string {
	return "download_stream_scraping"
}

func (s *CollectStreamObject) JobTime() string {
	return "0 */2 * * * *" // Tambah per 3 jam
}

func (s *CollectStreamObject) JobMiddleware(job *work.Job, next work.NextMiddlewareFunc) error {
	return next()
}

func (s *CollectStreamObject) Retry() uint {
	return 3
}

// handle download object scraping
func (s *CollectStreamObject) JobFunc(w *work.Job) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(60)*time.Second)
	defer cancel()

	err := s.streamProvider.DownloadStream(ctx)
	if err != nil {
		return err.Err[0]
	}

	return nil
}

type CollectStreamObjectUpdate struct {
	streamProvider provider.AllProviderStreaming
}

func NewCollectStreamObjectUpdate(provider provider.AllProviderStreaming) *CollectStreamObjectUpdate {
	return &CollectStreamObjectUpdate{streamProvider: provider}
}

func (s *CollectStreamObjectUpdate) JobName() string {
	return "download_stream_update"
}

func (s *CollectStreamObjectUpdate) JobTime() string {
	return "0 */3 * * * *"
}

func (s *CollectStreamObjectUpdate) JobMiddleware(job *work.Job, next work.NextMiddlewareFunc) error {
	return next()
}

func (s *CollectStreamObjectUpdate) Retry() uint {
	return 3
}

func (s *CollectStreamObjectUpdate) JobFunc(w *work.Job) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(60)*time.Second)
	defer cancel()

	err := s.streamProvider.DownloadStreamUpdate(ctx)
	if err != nil {
		return err.Err[0]
	}

	return nil
}
