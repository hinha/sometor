package usecase

import (
	"context"
	"github.com/hinha/sometor/entity"
	"github.com/hinha/sometor/provider"
)

type ListAccountTwitterOauth struct{}

func (l *ListAccountTwitterOauth) Perform(ctx context.Context, UserID string, provider provider.StreamKeyword) ([]entity.OUserTwitterInfo, *entity.ApplicationError) {
	users, err := provider.FindAllOauthTwitter(ctx, UserID)
	if err != nil {
		return users, err
	}

	return users, nil
}
