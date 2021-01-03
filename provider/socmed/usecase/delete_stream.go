package usecase

import (
	"context"
	"errors"
	"github.com/hinha/sometor/entity"
	"github.com/hinha/sometor/provider"
	"net/http"
)

type DeleteStream struct{}

func (d *DeleteStream) Perform(ctx context.Context, ID int, userID string, providerKeyword provider.StreamKeyword) (int, *entity.ApplicationError) {
	result, errProvider := providerKeyword.DeleteKeywordStream(ctx, ID, userID)
	if errProvider != nil {
		return 0, errProvider
	}

	if result == 0 {
		return 0, &entity.ApplicationError{
			Err:        []error{errors.New("keyword not found")},
			HTTPStatus: http.StatusNotFound,
		}
	}

	return result, nil
}
