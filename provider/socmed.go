package provider

import (
	"context"
	"github.com/hinha/sometor/entity"
)

type SocmedKeywordAPI interface {
	StreamKeywordList(ctx context.Context, ID string) ([]entity.StreamSequenceInitTable, *entity.ApplicationError)
	StreamKeywordCreate(ctx context.Context, request entity.StreamSequenceInsertable) (entity.StreamSequenceInsertable, *entity.ApplicationError)
	StreamKeywordDelete(ctx context.Context, ID int, userID string) (int, *entity.ApplicationError)
	StreamKeywordShowDataTwitter(ctx context.Context, media string, ID string, Keyword string) (entity.TwitterResult, *entity.ApplicationError)
	StreamKeywordShowDataInstagram(ctx context.Context, media string, ID string, Keyword string) (entity.InstagramResult, *entity.ApplicationError)
	StreamKeywordUpdateDataTwitter(ctx context.Context, media string, ID string, Keyword string) (entity.TwitterResult, *entity.ApplicationError)
	StreamKeywordUpdateDataInstagram(ctx context.Context, media string, ID string, Keyword string) (entity.InstagramResult, *entity.ApplicationError)
	StreamKeywordShowDataFacebook(ctx context.Context, media, ID, keyword string) (entity.FacebookResult, *entity.ApplicationError)
	TwitterOauthToken(ctx context.Context, request entity.OUserTwitter) (entity.OUserTwitterInfo, *entity.ApplicationError)
	TwitterListOauthAccount(ctx context.Context, userID string) ([]entity.OUserTwitterInfo, *entity.ApplicationError)
	TwitterPostFeedOauth(ctx context.Context, request entity.OFeedTwitter) *entity.ApplicationError
	TwitterListStatuses(ctx context.Context, UserTweetID, UserID string) ([]entity.OFeedTwitterInfo, *entity.ApplicationError)
	TwitterDelAccountOauth(ctx context.Context, UserTweetID, UserID string) (int, *entity.ApplicationError)
	TwitterPostFeedAllOauth(ctx context.Context, request entity.OFeedTwitterAll) *entity.ApplicationError
}
