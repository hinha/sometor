package scheduler

import (
	"context"
	"github.com/hinha/sometor/entity"
	"github.com/hinha/sometor/provider"
	"github.com/hinha/sometor/provider/scheduler/job"
	"github.com/hinha/sometor/provider/scheduler/usecase"
)

type StreamKeywordLocal struct {
	userProvider   provider.StreamSequence
	celeryProvider provider.CeleryClient
	s3Provider     provider.S3Management
}

func FabricateKeyword(provider provider.StreamSequence, celery provider.CeleryClient, s3 provider.S3Management) *StreamKeywordLocal {
	return &StreamKeywordLocal{userProvider: provider, celeryProvider: celery, s3Provider: s3}
}

func (s *StreamKeywordLocal) FabricateSchedule(engine provider.ScheduleEngine) {
	engine.Inject(job.NewSequenceAccount(s))
}

func (s *StreamKeywordLocal) CollectAccount(ctx context.Context) *entity.ApplicationError {
	find := usecase.FindCollectionAccount{}
	return find.PerformCollection(ctx, s.userProvider, s.celeryProvider, s.s3Provider)
}

func (s *StreamKeywordLocal) DownloadStream(ctx context.Context) *entity.ApplicationError {
	find := usecase.FindObjectS3Job{}
	return find.PerformCollection(ctx, s.userProvider, s.s3Provider)
}
