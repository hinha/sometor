package usecase

import (
	"context"
	"github.com/hinha/sometor/entity"
	"github.com/hinha/sometor/provider"
	"gopkg.in/validator.v2"
	"net/http"
)

type ProviderOauthTwitter struct{}

func (p *ProviderOauthTwitter) Perform(ctx context.Context, request entity.OUserTwitter, provider provider.StreamKeyword) (entity.OUserTwitterInfo, *entity.ApplicationError) {
	if err := p.Validate(request); err != nil {
		return entity.OUserTwitterInfo{}, &entity.ApplicationError{
			Err:        []error{err},
			HTTPStatus: http.StatusBadRequest,
		}
	}

	user, err := provider.CreateOrFindOauthTwitter(ctx, request)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (p *ProviderOauthTwitter) Validate(request entity.OUserTwitter) error {
	if err := validator.Validate(request); err != nil {
		return err
	}
	return nil
}
