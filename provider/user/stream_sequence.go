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

func (s *StreamSequence) FindAllUserMedia(ctx context.Context, media string) ([]entity.StreamSequenceInitTable, *entity.ApplicationError) {
	findAllByMedia := usecase.StreamRequestMedia{}
	return findAllByMedia.Perform(ctx, media, s.db)
}

func (s *StreamSequence) FindByUserAccountIDInfo(ctx context.Context, ID string) (entity.UserAccountSelectable, *entity.ApplicationError) {
	findUser := usecase.UserAccountID{}
	return findUser.PerformInfo(ctx, ID, s.db)
}

func (s *StreamSequence) FindByKeywordStreamInfo(ctx context.Context, keyword, media string) (entity.StreamSequenceInitTable, *entity.ApplicationError) {
	findKeyword := usecase.FindByKeywordStream{}
	return findKeyword.Perform(ctx, keyword, media, s.db)
}
