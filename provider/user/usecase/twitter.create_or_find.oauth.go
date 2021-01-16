package usecase

import (
	"context"
	"github.com/hinha/sometor/entity"
	"github.com/hinha/sometor/provider"
)

type CreateOrFindOauthTw struct{}

func (c *CreateOrFindOauthTw) Perform(ctx context.Context, request entity.OUserTwitter, streamProvider provider.StreamKeyword) (entity.OUserTwitterInfo, *entity.ApplicationError) {
	user, err := streamProvider.FindIdOauthTwitter(ctx, request.UserID)
	if err != nil && err.Err[0].Error() == "user not found" {
		err := streamProvider.CreateOauthTwitter(ctx, request)
		if err != nil {
			return user, err
		}

		user, err := streamProvider.FindIdOauthTwitter(ctx, request.UserID)
		if err != nil {
			return user, err
		}

		return user, nil
	} else if err != nil {
		return user, err
	}

	return user, nil
}
