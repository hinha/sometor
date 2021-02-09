package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hinha/sometor/entity"
	"github.com/hinha/sometor/provider"
	"io/ioutil"
	"net/http"
)

type ShowStreamTwitter struct{}

func (s *ShowStreamTwitter) Perform(ctx context.Context, media, userId, userKeyword string, provider provider.StreamKeyword) (entity.TwitterResult, *entity.ApplicationError) {
	_, errProvider := provider.FindByKeywordStreamWithAccount(ctx, userKeyword, media, "account", userId)
	if errProvider != nil {
		return entity.TwitterResult{}, errProvider
	}

	formatFile := fmt.Sprintf("temp/account-%s-%s.json", userKeyword, media)
	p, err := ioutil.ReadFile(formatFile)
	if err != nil {
		return entity.TwitterResult{}, &entity.ApplicationError{
			Err:        []error{err},
			HTTPStatus: http.StatusOK,
		}
	}

	var data entity.TwitterResult
	json.Unmarshal(p, &data)

	return data, nil
}

type ShowStreamInstagram struct{}

func (s *ShowStreamInstagram) Perform(ctx context.Context, media, userId, userKeyword string, provider provider.StreamKeyword) (entity.InstagramResult, *entity.ApplicationError) {
	_, errProvider := provider.FindByKeywordStreamWithAccount(ctx, userKeyword, media, "account", userId)
	if errProvider != nil {
		return entity.InstagramResult{}, errProvider
	}

	formatFile := fmt.Sprintf("temp/account-%s-%s.json", userKeyword, media)
	p, err := ioutil.ReadFile(formatFile)
	if err != nil {
		return entity.InstagramResult{}, &entity.ApplicationError{
			Err:        []error{err},
			HTTPStatus: http.StatusOK,
		}
	}

	var data entity.InstagramResult
	json.Unmarshal(p, &data)

	return data, nil
}
