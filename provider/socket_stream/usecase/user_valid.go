package usecase

import (
	"context"
	"github.com/hinha/sometor/entity"
	"github.com/hinha/sometor/provider"
)

type UserValidSocket struct{}

func (u *UserValidSocket) Perform(ctx context.Context, ID, fileName, media string, userProvider provider.StreamSequence) (entity.UserAccountSelectable, *entity.ApplicationError) {
	account, err := userProvider.FindByUserAccountIDInfo(ctx, ID)
	if err != nil {
		return entity.UserAccountSelectable{}, err
	}
	_, err = userProvider.FindByKeywordStreamInfo(ctx, fileName, media)
	if err != nil {
		return entity.UserAccountSelectable{}, err
	}

	return account, nil
}
