package user

import (
	"context"
	"github.com/hinha/sometor/entity"
	"github.com/hinha/sometor/provider"
	"github.com/hinha/sometor/provider/user/usecase"
)

type StreamKeyword struct {
	db provider.DB
}

func FabricateStreamKeyword(db provider.DB) *StreamKeyword {
	return &StreamKeyword{db: db}
}

func (s *StreamKeyword) FindByUserAccountIDInfo(ctx context.Context, ID string) (entity.UserAccountSelectable, *entity.ApplicationError) {
	findUser := usecase.UserAccountID{}
	return findUser.PerformInfo(ctx, ID, s.db)
}

func (s *StreamKeyword) FindAllStreamKeyword(ctx context.Context, ID string) ([]entity.StreamSequenceInitTable, *entity.ApplicationError) {
	findAll := usecase.StreamRequestID{}
	return findAll.Perform(ctx, ID, s.db)
}

func (s *StreamKeyword) FindByKeywordStreamWithAccount(ctx context.Context, keyword, media, types, userID string) (entity.StreamSequenceInsertable, *entity.ApplicationError) {
	findBy := usecase.FindByKeywordStream{}
	return findBy.PerformWithAccount(ctx, keyword, media, types, userID, s.db)
}

func (s *StreamKeyword) FindStreamKeywordID(ctx context.Context, ID int) (entity.StreamSequenceInsertable, *entity.ApplicationError) {
	findID := usecase.FindStreamKeywordID{}
	return findID.Perform(ctx, ID, s.db)
}

func (s *StreamKeyword) CreateOrFindStreamKeyword(ctx context.Context, request entity.StreamSequenceInsertable) (entity.StreamSequenceInsertable, *entity.ApplicationError) {
	createOrFind := usecase.CreateOrFindKeywordStream{}
	return createOrFind.Perform(ctx, request, s)
}

func (s *StreamKeyword) CreateKeywordStream(ctx context.Context, request entity.StreamSequenceInsertable) (int, *entity.ApplicationError) {
	create := usecase.CreateKeywordStream{}
	return create.Perform(ctx, request, s.db)
}

func (s *StreamKeyword) DeleteKeywordStream(ctx context.Context, ID int, userID string) (int, *entity.ApplicationError) {
	deleteKeywordStream := usecase.DeleteKeywordStream{}
	return deleteKeywordStream.Perform(ctx, ID, userID, s.db)
}

func (s *StreamKeyword) CreateOauthTwitter(ctx context.Context, request entity.OUserTwitter) *entity.ApplicationError {
	createAuth := usecase.CreateOauthTw{}
	return createAuth.Perform(ctx, request, s.db)
}

func (s *StreamKeyword) FindIdOauthTwitter(ctx context.Context, Id string) (entity.OUserTwitterInfo, *entity.ApplicationError) {
	findID := usecase.FindByIDOauthTw{}
	return findID.Perform(ctx, Id, s.db)
}

func (s *StreamKeyword) CreateOrFindOauthTwitter(ctx context.Context, request entity.OUserTwitter) (entity.OUserTwitterInfo, *entity.ApplicationError) {
	createOfFind := usecase.CreateOrFindOauthTw{}
	return createOfFind.Perform(ctx, request, s)
}
