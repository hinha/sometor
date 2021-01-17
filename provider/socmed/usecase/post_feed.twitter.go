package usecase

import (
	"context"
	twAPI "github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/hinha/sometor/entity"
	"github.com/hinha/sometor/provider"
	"net/http"
	"os"
)

type PostFeedTwitterOauth struct{}

func (p *PostFeedTwitterOauth) Perform(ctx context.Context, request entity.OFeedTwitter, providerKeyword provider.StreamKeyword) *entity.ApplicationError {
	users, err := providerKeyword.FindIDKeyOauthTwitter(ctx, request.UserTweetID, request.UserID)
	if err != nil {
		return err
	}

	config := oauth1.NewConfig(os.Getenv("TWITTER_CONSUMER_KEY"), os.Getenv("TWITTER_CONSUMER_SECRET"))
	token := oauth1.NewToken(users.AccessToken, users.AccessTokenSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twAPI.NewClient(httpClient)

	tweet, _, errs := client.Statuses.Update(request.Text, nil)
	if errs != nil {
		return &entity.ApplicationError{
			Err:        []error{errs},
			HTTPStatus: http.StatusBadRequest,
		}
	}

	err = providerKeyword.CreateTweetPostOauth(ctx, request, tweet.IDStr, tweet.User.ScreenName)
	if err != nil {
		return err
	}

	return nil
}