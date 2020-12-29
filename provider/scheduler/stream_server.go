package scheduler

import (
	"context"
	"github.com/hinha/sometor/entity"
	"github.com/hinha/sometor/provider"
	"github.com/hinha/sometor/provider/scheduler/job"
	"github.com/hinha/sometor/provider/scheduler/usecase"
)

type StreamKeywordServer struct {
	userProvider   provider.StreamSequence
	celeryProvider provider.CeleryClient
	s3Provider     provider.S3Management
}

func FabricateKeywordServer(provider provider.StreamSequence, celery provider.CeleryClient, s3 provider.S3Management) *StreamKeywordServer {
	return &StreamKeywordServer{userProvider: provider, celeryProvider: celery, s3Provider: s3}
}

func (s *StreamKeywordServer) FabricateSchedule(engine provider.ScheduleEngine) {
	engine.Inject(job.NewCollectStream(s))
	engine.Inject(job.NewCollectStreamObjectUpdate(s))
}

func (s *StreamKeywordServer) CollectAccount(ctx context.Context) *entity.ApplicationError {
	find := usecase.FindCollectionAccount{}
	return find.PerformCollection(ctx, s.userProvider, s.celeryProvider, s.s3Provider)
}

func (s *StreamKeywordServer) DownloadStream(ctx context.Context) *entity.ApplicationError {
	find := usecase.FindObjectS3Job{}
	return find.PerformCollection(ctx, s.userProvider, s.s3Provider)
}

func (s *StreamKeywordServer) DownloadStreamUpdate(ctx context.Context) *entity.ApplicationError {
	find := usecase.FindObjectS3Job{}
	return find.PerformCollectionUpdate(ctx, s.userProvider, s.s3Provider)
}
