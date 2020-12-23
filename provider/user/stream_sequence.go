package user

import (
	"context"
	"github.com/hinha/sometor/entity"
	"github.com/hinha/sometor/provider"
	"github.com/hinha/sometor/provider/user/usecase"
)

type StreamSequence struct {
	db provider.DB
}

func FabricateStream(db provider.DB) *StreamSequence {
	return &StreamSequence{db: db}
}

func (s *StreamSequence) FindByUserID(ctx context.Context, ID string) (entity.StreamSequenceInitTable, *entity.ApplicationError) {
	find := usecase.StreamRequestByID{}
	return find.Perform(ctx, ID, s.db)
}

func (s *StreamSequence) FindAllUser(ctx context.Context) ([]entity.StreamSequenceInitTable, *entity.ApplicationError) {
	findAll := usecase.StreamRequestAll{}
	return findAll.Perform(ctx, s.db)
}
