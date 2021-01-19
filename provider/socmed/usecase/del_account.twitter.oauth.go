package usecase

import (
	"context"
	"errors"
	"github.com/hinha/sometor/entity"
	"github.com/hinha/sometor/provider"
	"net/http"
)

type DelAccountTwitterOauth struct{}

func (d *DelAccountTwitterOauth) Perform(ctx context.Context, tweetUserID, userID string, provider provider.StreamKeyword) (int, *entity.ApplicationError) {
	result, err := provider.DeleteTweetAccountOauth(ctx, tweetUserID, userID)
	if err != nil {
		return 0, err
	}

	if result == 0 {
		return 0, &entity.ApplicationError{
			Err:        []error{errors.New("tweet account not found")},
			HTTPStatus: http.StatusNotFound,
		}
	}

	return result, nil
}
