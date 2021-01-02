package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/bradfitz/slice"
	"github.com/hinha/sometor/entity"
	"github.com/hinha/sometor/provider"
	"io/ioutil"
)

type UpdateStreamTwitter struct{}

func (u *UpdateStreamTwitter) Perform(ctx context.Context, media, userId, userKeyword string, provider provider.StreamKeyword) (entity.TwitterResult, *entity.ApplicationError) {
	_, errProvider := provider.FindByKeywordStreamWithAccount(ctx, userKeyword, media, "account", userId)
	if errProvider != nil {
		return entity.TwitterResult{}, errProvider
	}

	formatFile := fmt.Sprintf("temp/update/account-%s-%s.json", userKeyword, media)
	p, err := ioutil.ReadFile(formatFile)
	if err != nil {
		return entity.TwitterResult{}, &entity.ApplicationError{
			Err: []error{err},
		}
	}

	var data entity.TwitterResult
	_ = json.Unmarshal(p, &data)

	if len(data.Results) > 0 {
		data.Results = data.Results[:1]
	}
	return data, nil
}

type UpdateStreamInstagram struct{}

func (u *UpdateStreamInstagram) Perform(ctx context.Context, media, userId, userKeyword string, provider provider.StreamKeyword) (entity.InstagramResult, *entity.ApplicationError) {
	_, errProvider := provider.FindByKeywordStreamWithAccount(ctx, userKeyword, media, "account", userId)
	if errProvider != nil {
		return entity.InstagramResult{}, errProvider
	}

	formatFile := fmt.Sprintf("temp/account-%s-%s.json", userKeyword, media)
	p, err := ioutil.ReadFile(formatFile)
	if err != nil {
		return entity.InstagramResult{}, &entity.ApplicationError{
			Err: []error{err},
		}
	}

	var data entity.InstagramResult
	_ = json.Unmarshal(p, &data)

	if len(data.Results) > 0 {
		slice.Sort(data.Results[:], func(i, j int) bool {
			return data.Results[i].Timestamp > data.Results[j].Timestamp
		})
		data.Results = data.Results[:1]
		data.LastUpdate = data.Results[0].CreatedAt
	}

	return data, nil
}
