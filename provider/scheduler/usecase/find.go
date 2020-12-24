package usecase

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hinha/sometor/entity"
	"github.com/hinha/sometor/provider"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"strings"
)

type Claims map[string]interface{}
type FindCollectionAccount struct{}

func (f *FindCollectionAccount) PerformCollection(ctx context.Context, userProvider provider.StreamSequence, celeryProvider provider.CeleryClient, s3Provider provider.S3Management) ([]entity.StreamSequenceInitTable, *entity.ApplicationError) {
	//var arrData []map[string]interface{}
	// perbanyak
	// 1. collect data from db
	// 2. call rpc twitter or instagram
	// 3. create file
	// 4. put object to s3

	// flow
	//  cron <- db -> infra task -> s3
	account, err := userProvider.FindAllUser(ctx)
	if err != nil {
		return account, err
	}

	//ID            int       `json:"id"`
	//	Keyword       string    `json:"keyword"`
	//	Media         string    `json:"media"`
	//	Type          string    `json:"type"`
	//	CreatedAt     time.Time `json:"created_at"`
	//	UserAccountID string    `json:"user_account_id"`
	var keyword string
	if collections, err := userProvider.FindAllUser(ctx); err != nil {
		return collections, err
	} else {
		for _, data := range collections {
			if data.Media == "twitter" {
				result, errCelery := celeryProvider.GetTaskResult("example.twitter_scrape_v1", 1, data)
				if errCelery != nil {
					return nil, &entity.ApplicationError{
						Err:        []error{errors.New("task error")},
						HTTPStatus: http.StatusNotFound,
					}
				}
				if result != nil {
					keyword = strings.ReplaceAll(data.Keyword, "@", "")
					osFile, pathString, err := f.CreateFileFromMap(keyword, data.Media, result)
					if err != nil {
						fmt.Println(err)
					}

					err = s3Provider.PutObject(pathString, osFile)
					if err != nil {
						fmt.Println(err)
						return nil, &entity.ApplicationError{
							Err:        []error{errors.New("cannot put object s3")},
							HTTPStatus: http.StatusNotFound,
						}
					}
				}

			} else if data.Media == "instagram" {
				fmt.Println("instagram Data: ", data)
			} else {
				fmt.Println("facebook Data: ", data)
			}
		}
	}

	return userProvider.FindAllUser(ctx)
}

func (f *FindCollectionAccount) FillStruct(m map[string]interface{}, s interface{}) error {
	structValue := reflect.ValueOf(s).Elem()

	for name, value := range m {
		structFieldValue := structValue.FieldByName(name)

		if !structFieldValue.IsValid() {
			return fmt.Errorf("No such field: %s in obj", name)
		}

		if !structFieldValue.CanSet() {
			return fmt.Errorf("Cannot set %s field value", name)
		}

		val := reflect.ValueOf(value)
		if structFieldValue.Type() != val.Type() {
			return errors.New("Provided value type didn't match obj field type")
		}

		structFieldValue.Set(val)
	}
	return nil
}

func (f *FindCollectionAccount) CreateFileFromMap(fileName, mediaType string, m interface{}) (*os.File, string, error) {

	data, _ := json.MarshalIndent(m, "", "  ")

	basePath, _ := os.Getwd()
	path := fmt.Sprintf("%s/temp", basePath)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			return nil, "", err
		}
	}

	saveFile := fmt.Sprintf("temp/%s-%s.json", fileName, mediaType)
	_ = ioutil.WriteFile(fmt.Sprintf("%s/%s", basePath, saveFile), data, os.ModePerm)

	file, err := os.Open(saveFile)
	if err != nil {
		return nil, "", errors.New(fmt.Sprintf("Unable to open file %v", err))
	}

	return file, saveFile, nil
}
