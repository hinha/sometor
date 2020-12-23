package scheduler

import (
	"context"
	"github.com/hinha/sometor/entity"
	"github.com/hinha/sometor/provider"
	job2 "github.com/hinha/sometor/provider/scheduler/job"
	"github.com/hinha/sometor/provider/scheduler/usecase"
)

type StreamKeyword struct {
	userProvider   provider.StreamSequence
	celeryProvider provider.CeleryClient
}

func FabricateKeyword(provider provider.StreamSequence, celery provider.CeleryClient) *StreamKeyword {
	return &StreamKeyword{userProvider: provider, celeryProvider: celery}
}

func (s *StreamKeyword) FabricateSchedule(engine provider.ScheduleEngine) {
	engine.Inject(job2.NewSequenceAccount(s))
}

func (s *StreamKeyword) CollectAccount(ctx context.Context) ([]entity.StreamSequenceInitTable, *entity.ApplicationError) {
	find := usecase.FindCollectionAccount{}
	return find.PerformCollection(ctx, s.userProvider, s.celeryProvider)
}
