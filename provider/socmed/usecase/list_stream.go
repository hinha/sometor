package usecase

import (
	"context"
	"fmt"
	"github.com/hinha/sometor/entity"
	"github.com/hinha/sometor/provider"
)

type ListStream struct {
}

func (l *ListStream) Perform(ctx context.Context, ID string, providerKeyword provider.StreamKeyword) ([]entity.StreamSequenceInitTable, *entity.ApplicationError) {
	keyword, err := providerKeyword.FindByUserAccountIDInfo(ctx, ID)
	if err != nil {
		return nil, err
	}
	fmt.Println(keyword)

	return providerKeyword.FindAllStreamKeyword(ctx, keyword.UniqueAccount)
}
