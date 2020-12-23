package usecase

import (
	"context"
	"fmt"
	"github.com/hinha/sometor/entity"
	"github.com/hinha/sometor/provider"
)

type FindCollectionAccount struct{}

func (f *FindCollectionAccount) PerformCollection(ctx context.Context, userProvider provider.StreamSequence) ([]entity.StreamSequenceInitTable, *entity.ApplicationError) {
	// perbanyak
	// 1. collect data from db
	// 2. call rpc twitter or instagram
	// 3. create file
	// 4. put object to s3

	// flow
	//  cron <- db -> infra rpc -> s3
	account, err := userProvider.FindAllUser(ctx)
	if err != nil {
		return account, err
	}

	if collections, err := userProvider.FindAllUser(ctx); err != nil {
		return collections, err
	} else {
		for _, data := range collections {
			if data.Media == "twitter" {
				fmt.Println("twitter Data: ", data)
			} else if data.Media == "instagram" {
				fmt.Println("instagram Data: ", data)
			} else {
				fmt.Println("facebook Data: ", data)
			}
		}
	}

	return userProvider.FindAllUser(ctx)
}
