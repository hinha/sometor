package provider

import (
	"context"
	"github.com/hinha/sometor/entity"
)

type StreamSequence interface {
	FindByUserID(ctx context.Context, ID string) (entity.StreamSequenceInitTable, *entity.ApplicationError)
	FindAllUser(ctx context.Context) ([]entity.StreamSequenceInitTable, *entity.ApplicationError)
	FindByUserAccountIDInfo(ctx context.Context, ID string) (entity.UserAccountSelectable, *entity.ApplicationError)
	FindByKeywordStreamInfo(ctx context.Context, keyword, media string) (entity.StreamSequenceInitTable, *entity.ApplicationError)
}

type StreamKeyword interface {
	FindByUserAccountIDInfo(ctx context.Context, ID string) (entity.UserAccountSelectable, *entity.ApplicationError)
	FindAllStreamKeyword(ctx context.Context, ID string) ([]entity.StreamSequenceInitTable, *entity.ApplicationError)
	FindByKeywordStreamWithAccount(ctx context.Context, keyword, media, types, userID string) (entity.StreamSequenceInsertable, *entity.ApplicationError)
	FindStreamKeywordID(ctx context.Context, ID int) (entity.StreamSequenceInsertable, *entity.ApplicationError)
	CreateKeywordStream(ctx context.Context, request entity.StreamSequenceInsertable) (int, *entity.ApplicationError)
	CreateOrFindStreamKeyword(ctx context.Context, request entity.StreamSequenceInsertable) (entity.StreamSequenceInsertable, *entity.ApplicationError)
	DeleteKeywordStream(ctx context.Context, ID int, userID string) (int, *entity.ApplicationError)
}
