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
}
