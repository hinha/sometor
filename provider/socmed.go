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
	TwitterOauthToken(ctx context.Context, request entity.OUserTwitter) (entity.OUserTwitterInfo, *entity.ApplicationError)
	TwitterListOauthAccount(ctx context.Context, userID string) ([]entity.OUserTwitterInfo, *entity.ApplicationError)
}
