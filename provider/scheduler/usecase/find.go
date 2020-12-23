package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/hinha/sometor/entity"
	"github.com/hinha/sometor/provider"
	"log"
	"math/rand"
	"net/http"
	"reflect"
)

type FindCollectionAccount struct{}

func (f *FindCollectionAccount) PerformCollection(ctx context.Context, userProvider provider.StreamSequence, celeryProvider provider.CeleryClient) ([]entity.StreamSequenceInitTable, *entity.ApplicationError) {
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

	// prepare arguments
	argA := rand.Intn(10)
	argB := rand.Intn(10)

	result, errCelery := celeryProvider.GetTaskResult("example.add", 1, argA, argB)
	if errCelery != nil {
		return nil, &entity.ApplicationError{
			Err:        []error{errors.New("task error")},
			HTTPStatus: http.StatusNotFound,
		}
	}
	log.Printf("result: %+v of type %+v", result, reflect.TypeOf(result))

	return userProvider.FindAllUser(ctx)
}
