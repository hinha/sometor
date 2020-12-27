package usecase

import (
	"context"
	"github.com/hinha/sometor/entity"
	"github.com/hinha/sometor/provider"
)

type CreateOrFindKeywordStream struct{}

func (c *CreateOrFindKeywordStream) Perform(ctx context.Context, request entity.StreamSequenceInsertable, streamProvider provider.StreamKeyword) (entity.StreamSequenceInsertable, *entity.ApplicationError) {
	keyword, err := streamProvider.FindByKeywordStreamWithAccount(ctx, request.Keyword, request.Media, request.Type, request.UserAccountID)
	if err != nil && err.Err[0].Error() == "keyword not found" {
		id, err := streamProvider.CreateKeywordStream(ctx, request)
		if err != nil {
			return keyword, err
		}

		keyword, err := streamProvider.FindStreamKeywordID(ctx, id)
		if err != nil {
			return keyword, err
		}

		return keyword, nil
	} else if err != nil {
		return keyword, err
	}

	return keyword, nil
}
