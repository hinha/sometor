package usecase

import (
	"context"
	twAPI "github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/hinha/sometor/entity"
	"github.com/hinha/sometor/provider"
	"os"
	"strconv"
)

type ListStatusesTweetOauth struct{}

func (l *ListStatusesTweetOauth) Perform(ctx context.Context, userTweetID, userId string, provider provider.StreamKeyword) ([]entity.OFeedTwitterInfo, *entity.ApplicationError) {

	users, err := provider.FindIDKeyOauthTwitter(ctx, userTweetID, userId)
	if err != nil {
		return []entity.OFeedTwitterInfo{}, err
	}

	config := oauth1.NewConfig(os.Getenv("TWITTER_CONSUMER_KEY"), os.Getenv("TWITTER_CONSUMER_SECRET"))
	token := oauth1.NewToken(users.AccessToken, users.AccessTokenSecret)
	httpClient := config.Client(oauth1.NoContext, token)
	// Twitter client
	client := twAPI.NewClient(httpClient)

	statuses, err := provider.FindAllStatusesTweetOauth(ctx, users.UserTweetID)
	if err != nil {
		return []entity.OFeedTwitterInfo{}, err
	}

	for i := 0; i < len(statuses); i++ {
		if len(statuses) != 0 {
			id, _ := strconv.Atoi(statuses[i].StatusesID)
			tweetPost, _, err := client.Statuses.Show(int64(id), nil)
			if err == nil {
				if tweetPost.ExtendedEntities != nil {
					for _, media := range tweetPost.ExtendedEntities.Media {
						statuses[i].ImageUrlHttps = media.MediaURLHttps
					}
				}
			}
		}
	}

	return statuses, nil
}
