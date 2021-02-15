package socmed

import (
	"context"
	"errors"
	"github.com/hinha/sometor/entity"
	"github.com/hinha/sometor/provider"
	"github.com/hinha/sometor/provider/socmed/api"
	"github.com/hinha/sometor/provider/socmed/usecase"
	"net/http"
	"strings"
)

type KeywordStream struct {
	providerKeyword provider.StreamKeyword
}

func FabricateKeyword(providerKeyword provider.StreamKeyword) *KeywordStream {
	return &KeywordStream{providerKeyword: providerKeyword}
}

func (k *KeywordStream) FabricateAPI(engine provider.APIEngine) {
	engine.InjectAPI(api.NewListStreamKeyword(k))
	engine.InjectAPI(api.NewCreateStreamKeyword(k))
	engine.InjectAPI(api.NewDeleteStreamKeyword(k))
	engine.InjectAPI(api.NewShowStreamData(k))
	engine.InjectAPI(api.NewUpdateStreamData(k))
	engine.InjectAPI(api.NewOauthCallback(k))
	engine.InjectAPI(api.NewListAccountTwitterOauth(k))
	engine.InjectAPI(api.NewPostFeedTwitter(k))
	engine.InjectAPI(api.NewListStatusesTwitterOauth(k))
	engine.InjectAPI(api.NewDelAccountTwitterOauth(k))
	engine.InjectAPI(api.NewPostFeedTwitterMulti(k))
	//engine.InjectAPI(api.NewOauthClientToken(k))
}

func (k *KeywordStream) StreamKeywordList(ctx context.Context, ID string) ([]entity.StreamSequenceInitTable, *entity.ApplicationError) {
	listKeyword := usecase.ListStream{}
	return listKeyword.Perform(ctx, ID, k.providerKeyword)
}

func (k *KeywordStream) StreamKeywordCreate(ctx context.Context, request entity.StreamSequenceInsertable) (entity.StreamSequenceInsertable, *entity.ApplicationError) {
	var mediaValid bool

	switch request.Media {
	case "twitter", "instagram", "facebook":
		mediaValid = true
		break
	default:
		mediaValid = false
		break
	}

	if mediaValid == false {
		return entity.StreamSequenceInsertable{}, &entity.ApplicationError{
			Err:        []error{errors.New("media not found")},
			HTTPStatus: http.StatusNotFound,
		}
	}

	if request.Type == "account" {
		request.Keyword = strings.ReplaceAll(request.Keyword, "@", "")
	} else if request.Type == "hashtag" {
		request.Keyword = strings.ReplaceAll(request.Keyword, "#", "")
	}

	result, err := k.providerKeyword.CreateOrFindStreamKeyword(ctx, request)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (k *KeywordStream) StreamKeywordDelete(ctx context.Context, ID int, userID string) (int, *entity.ApplicationError) {
	deleteKeyword := usecase.DeleteStream{}
	return deleteKeyword.Perform(ctx, ID, userID, k.providerKeyword)
}

func (k *KeywordStream) StreamKeywordShowDataTwitter(ctx context.Context, media string, ID string, Keyword string) (entity.TwitterResult, *entity.ApplicationError) {
	streamTwitter := usecase.ShowStreamTwitter{}
	return streamTwitter.Perform(ctx, media, ID, Keyword, k.providerKeyword)
}

func (k *KeywordStream) StreamKeywordShowDataInstagram(ctx context.Context, media string, ID string, Keyword string) (entity.InstagramResult, *entity.ApplicationError) {
	streamInstagram := usecase.ShowStreamInstagram{}
	return streamInstagram.Perform(ctx, media, ID, Keyword, k.providerKeyword)
}

func (k *KeywordStream) StreamKeywordUpdateDataTwitter(ctx context.Context, media string, ID string, Keyword string) (entity.TwitterResult, *entity.ApplicationError) {
	streamTwitter := usecase.UpdateStreamTwitter{}
	return streamTwitter.Perform(ctx, media, ID, Keyword, k.providerKeyword)
}

func (k *KeywordStream) StreamKeywordUpdateDataInstagram(ctx context.Context, media string, ID string, Keyword string) (entity.InstagramResult, *entity.ApplicationError) {
	streamInstagram := usecase.UpdateStreamInstagram{}
	return streamInstagram.Perform(ctx, media, ID, Keyword, k.providerKeyword)
}

func (k *KeywordStream) TwitterOauthToken(ctx context.Context, request entity.OUserTwitter) (entity.OUserTwitterInfo, *entity.ApplicationError) {
	oauth := usecase.ProviderOauthTwitter{}
	return oauth.Perform(ctx, request, k.providerKeyword)
}

func (k *KeywordStream) TwitterListOauthAccount(ctx context.Context, userID string) ([]entity.OUserTwitterInfo, *entity.ApplicationError) {
	account := usecase.ListAccountTwitterOauth{}
	return account.Perform(ctx, userID, k.providerKeyword)
}

func (k *KeywordStream) TwitterPostFeedOauth(ctx context.Context, request entity.OFeedTwitter) *entity.ApplicationError {
	postFeed := usecase.PostFeedTwitterOauth{}
	return postFeed.Perform(ctx, request, k.providerKeyword)
}

func (k *KeywordStream) TwitterListStatuses(ctx context.Context, UserTweetID, UserID string) ([]entity.OFeedTwitterInfo, *entity.ApplicationError) {
	statuesList := usecase.ListStatusesTweetOauth{}
	return statuesList.Perform(ctx, UserTweetID, UserID, k.providerKeyword)
}

func (k *KeywordStream) TwitterDelAccountOauth(ctx context.Context, UserTweetID, UserID string) (int, *entity.ApplicationError) {
	delAccount := usecase.DelAccountTwitterOauth{}
	return delAccount.Perform(ctx, UserTweetID, UserID, k.providerKeyword)
}

func (k *KeywordStream) TwitterPostFeedAllOauth(ctx context.Context, request entity.OFeedTwitterAll) *entity.ApplicationError {
	if request.ImageBase64 == "" {
		postFeed := usecase.PostFeedTwitterOauthMulti{}
		return postFeed.Perform(ctx, request, k.providerKeyword)
	} else {
		postFeedImg := usecase.PostImageFeedTweetMulti{}
		return postFeedImg.Perform(ctx, request, k.providerKeyword)
	}
}
