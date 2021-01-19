package usecase

import (
	"context"
	"errors"
	"github.com/hinha/sometor/entity"
	"github.com/hinha/sometor/provider"
	"net/http"
)

type FindAllStatusesOauthTweet struct{}

func (f *FindAllStatusesOauthTweet) Perform(ctx context.Context, userTweetID string, db provider.DB) ([]entity.OFeedTwitterInfo, *entity.ApplicationError) {
	var data []entity.OFeedTwitterInfo

	rows, err := db.QueryContext(ctx, "find-oauth-statuses", "select statuses_id,text,permalink,created_at from twitter_statuses where user_id = ?", userTweetID)
	if err != nil {
		return data, &entity.ApplicationError{
			Err:        []error{errors.New("twitter statuses not found")},
			HTTPStatus: http.StatusNotFound,
		}
	}

	for rows.Next() {
		var each entity.OFeedTwitterInfo
		if err := rows.Scan(&each.StatusesID, &each.Text, &each.Permalink, &each.CreatedAt); err != nil {
			return data, &entity.ApplicationError{
				Err:        []error{errors.New("service unavailable")},
				HTTPStatus: http.StatusServiceUnavailable,
			}
		}

		data = append(data, each)
	}

	if err = rows.Err(); err != nil {
		return data, &entity.ApplicationError{
			Err:        []error{errors.New("[rows] service unavailable")},
			HTTPStatus: http.StatusServiceUnavailable,
		}
	}

	return data, nil
}
