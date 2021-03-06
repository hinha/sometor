package provider

import (
	"context"
	"github.com/hinha/sometor/entity"
)

type StreamSequence interface {
	FindByUserID(ctx context.Context, ID string) (entity.StreamSequenceInitTable, *entity.ApplicationError)
	FindAllUser(ctx context.Context) ([]entity.StreamSequenceInitTable, *entity.ApplicationError) // deprecated
	FindAllUserMedia(ctx context.Context, media string) ([]entity.StreamSequenceInitTable, *entity.ApplicationError)
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
	CreateOauthTwitter(ctx context.Context, request entity.OUserTwitter) *entity.ApplicationError
	FindIdOauthTwitter(ctx context.Context, Id string) (entity.OUserTwitterInfo, *entity.ApplicationError)
	FindIDKeyOauthTwitter(ctx context.Context, Id, userID string) (entity.OUserTwitterKey, *entity.ApplicationError)
	FindAllIDKeyOauthTwitter(ctx context.Context, userID string) ([]entity.OUserTwitterKey, *entity.ApplicationError)
	CreateOrFindOauthTwitter(ctx context.Context, request entity.OUserTwitter) (entity.OUserTwitterInfo, *entity.ApplicationError)
	FindAllOauthTwitter(ctx context.Context, userID string) ([]entity.OUserTwitterInfo, *entity.ApplicationError)
	CreateTweetPostOauth(ctx context.Context, request entity.OFeedTwitter, statusID, username string) *entity.ApplicationError
	FindAllStatusesTweetOauth(ctx context.Context, userTweetID string) ([]entity.OFeedTwitterInfo, *entity.ApplicationError)
	DeleteTweetAccountOauth(ctx context.Context, userTweetID, userID string) (int, *entity.ApplicationError)
}
