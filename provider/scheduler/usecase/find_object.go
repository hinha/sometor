package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/hinha/sometor/entity"
	"github.com/hinha/sometor/provider"
	"strings"
)

type FindObjectS3Job struct{}

func (f *FindObjectS3Job) PerformCollection(ctx context.Context, userProvider provider.StreamSequence, s3Provider provider.S3Management) *entity.ApplicationError {
	var keyword string
	if collections, err := userProvider.FindAllUser(ctx); err != nil {
		return err
	} else {
		for _, data := range collections {
			switch data.Type {
			case "account":
				keyword = strings.ReplaceAll(data.Keyword, "@", "")
			case "hashtag":
				keyword = strings.ReplaceAll(data.Keyword, "#", "")
			}

			formatName := fmt.Sprintf("%s-%s-%s.json", data.Type, keyword, data.Media)
			if data.Media == "twitter" {
				_, err := s3Provider.DownloadObject(fmt.Sprintf("temp/%s", formatName))
				if err != nil {
					return &entity.ApplicationError{
						Err: []error{errors.New("cannot download twitter object s3")},
					}
				}
			} else if data.Media == "instagram" {
				_, err := s3Provider.DownloadObject(fmt.Sprintf("temp/%s", formatName))
				if err != nil {
					return &entity.ApplicationError{
						Err: []error{errors.New("cannot download instagram object s3")},
					}
				}

			} else {
				fmt.Println("facebook Data: ", data)
			}
		}
	}

	return nil
}
