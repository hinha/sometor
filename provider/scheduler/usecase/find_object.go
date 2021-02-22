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

	if collections, err := userProvider.FindAllUserMedia(ctx, "twitter"); err != nil {
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
			_, err := s3Provider.DownloadObject(fmt.Sprintf("temp/%s", formatName))
			if err != nil {
				return &entity.ApplicationError{
					Err: []error{errors.New("cannot download twitter object s3")},
				}
			}
		}
	}

	if collections, err := userProvider.FindAllUserMedia(ctx, "facebook"); err != nil {
		return err
	} else {
		for _, data := range collections {
			formatName := fmt.Sprintf("%s-%s-%s.json", data.Type, data.Keyword, data.Media)
			_, err := s3Provider.DownloadObject(fmt.Sprintf("temp/%s", formatName))
			if err != nil {
				return &entity.ApplicationError{
					Err: []error{errors.New("cannot download facebook object s3")},
				}
			}
		}
	}

	if collections, err := userProvider.FindAllUserMedia(ctx, "instagram"); err != nil {
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
			_, err := s3Provider.DownloadObject(fmt.Sprintf("temp/%s", formatName))
			if err != nil {
				return &entity.ApplicationError{
					Err: []error{errors.New("cannot download twitter object s3")},
				}
			}
		}
	}

	return nil
}

func (f *FindObjectS3Job) PerformCollectionUpdate(ctx context.Context, userProvider provider.StreamSequence, s3Provider provider.S3Management) *entity.ApplicationError {
	var keyword string

	if collections, err := userProvider.FindAllUserMedia(ctx, "twitter"); err != nil {
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
			_, err := s3Provider.DownloadObjectUpdate(fmt.Sprintf("temp/%s", formatName), formatName)
			if err != nil {
				return &entity.ApplicationError{
					Err: []error{errors.New("cannot download twitter object s3")},
				}
			}
		}
	}

	if collections, err := userProvider.FindAllUserMedia(ctx, "facebook"); err != nil {
		return err
	} else {
		for _, data := range collections {
			formatName := fmt.Sprintf("%s-%s-%s.json", data.Type, data.Keyword, data.Media)
			_, err := s3Provider.DownloadObjectUpdate(fmt.Sprintf("temp/%s", formatName), formatName)
			if err != nil {
				return &entity.ApplicationError{
					Err: []error{errors.New("cannot download facebook object s3")},
				}
			}
		}
	}

	if collections, err := userProvider.FindAllUserMedia(ctx, "instagram"); err != nil {
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
			_, err := s3Provider.DownloadObjectUpdate(fmt.Sprintf("temp/%s", formatName), formatName)
			if err != nil {
				return &entity.ApplicationError{
					Err: []error{errors.New("cannot download twitter object s3")},
				}
			}
		}
	}

	return nil
}
