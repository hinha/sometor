package usecase

import (
	"context"
	"github.com/hinha/sometor/entity"
	"github.com/hinha/sometor/provider"
)

type ListStatusesTweetOauth struct{}

func (l *ListStatusesTweetOauth) Perform(ctx context.Context, userTweetID, userId string, provider provider.StreamKeyword) ([]entity.OFeedTwitterInfo, *entity.ApplicationError) {
	users, err := provider.FindIDKeyOauthTwitter(ctx, userTweetID, userId)
	if err != nil {
		return []entity.OFeedTwitterInfo{}, err
	}

	statuses, err := provider.FindAllStatusesTweetOauth(ctx, users.UserTweetID)
	if err != nil {
		return []entity.OFeedTwitterInfo{}, err
	}

	return statuses, nil
}
