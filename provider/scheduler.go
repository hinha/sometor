package provider

import (
	"context"
	"github.com/gocraft/work"
	"github.com/hinha/sometor/entity"
)

type ScheduleHandler interface {
	JobFunc(context *work.Job) error
	JobName() string
	JobTime() string
	JobMiddleware(job *work.Job, next work.NextMiddlewareFunc) error
	Retry() uint
}

type ScheduleEngine interface {
	Run()
	Inject(handler ScheduleHandler)
	Shutdown(ctx context.Context)
}

type AllProviderStreaming interface {
	CollectAccount(ctx context.Context) *entity.ApplicationError       // Local
	DownloadStream(ctx context.Context) *entity.ApplicationError       // Server
	DownloadStreamUpdate(ctx context.Context) *entity.ApplicationError // Server
}
