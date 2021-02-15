package usecase

import (
	"context"
	"errors"
	"github.com/chimeracoder/anaconda"
	twAPI "github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/hinha/sometor/entity"
	"github.com/hinha/sometor/provider"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

type PostFeedTwitterOauth struct{}

func (p *PostFeedTwitterOauth) Perform(ctx context.Context, request entity.OFeedTwitter, providerKeyword provider.StreamKeyword) *entity.ApplicationError {
	users, err := providerKeyword.FindIDKeyOauthTwitter(ctx, request.UserTweetID, request.UserID)
	if err != nil {
		return err
	}

	config := oauth1.NewConfig(os.Getenv("TWITTER_CONSUMER_KEY"), os.Getenv("TWITTER_CONSUMER_SECRET"))
	token := oauth1.NewToken(users.AccessToken, users.AccessTokenSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twAPI.NewClient(httpClient)

	tweet, _, errs := client.Statuses.Update(request.Text, nil)
	if errs != nil {
		return &entity.ApplicationError{
			Err:        []error{errs},
			HTTPStatus: http.StatusBadRequest,
		}
	}

	err = providerKeyword.CreateTweetPostOauth(ctx, request, tweet.IDStr, tweet.User.ScreenName)
	if err != nil {
		return err
	}

	return nil
}

type PostFeedTwitterOauthMulti struct{}

func (p *PostFeedTwitterOauthMulti) Perform(ctx context.Context, request entity.OFeedTwitterAll, providerKeyword provider.StreamKeyword) *entity.ApplicationError {
	users, err := providerKeyword.FindAllIDKeyOauthTwitter(ctx, request.UserID)
	if err != nil {
		return err
	}

	config := oauth1.NewConfig(os.Getenv("TWITTER_CONSUMER_KEY"), os.Getenv("TWITTER_CONSUMER_SECRET"))
	for _, key := range users {
		token := oauth1.NewToken(key.AccessToken, key.AccessTokenSecret)
		httpClient := config.Client(oauth1.NoContext, token)

		// Twitter client
		client := twAPI.NewClient(httpClient)

		tweet, _, errs := client.Statuses.Update(request.Text, nil)
		if errs != nil {
			return &entity.ApplicationError{
				Err:        []error{errs},
				HTTPStatus: http.StatusBadRequest,
			}
		}

		requestCreate := entity.OFeedTwitter{
			UserID:      request.UserID,
			Text:        request.Text,
			UserTweetID: key.UserTweetID,
		}

		err = providerKeyword.CreateTweetPostOauth(ctx, requestCreate, tweet.IDStr, tweet.User.ScreenName)
		if err != nil {
			return err
		}
	}

	return nil
}

type PostImageFeedTweetMulti struct{}

func (p *PostImageFeedTweetMulti) Perform(ctx context.Context, request entity.OFeedTwitterAll, providerKeyword provider.StreamKeyword) *entity.ApplicationError {
	users, err := providerKeyword.FindAllIDKeyOauthTwitter(ctx, request.UserID)
	if err != nil {
		return err
	}

	if request.ImageBase64 == "" {
		return &entity.ApplicationError{
			Err:        []error{errors.New("image must required")},
			HTTPStatus: http.StatusBadRequest,
		}
	} else if request.Text == "" {
		return &entity.ApplicationError{
			Err:        []error{errors.New("caption must required")},
			HTTPStatus: http.StatusBadRequest,
		}
	}

	var validTypeImg bool
	colImg := strings.Split(request.ImageBase64, ",")
	switch strings.TrimSuffix(request.ImageBase64[5:len(colImg[0])], ";base64") {
	case "image/png":
		validTypeImg = true
	case "image/jpeg":
		validTypeImg = true
	case "image/jpg":
		validTypeImg = true
	default:
		validTypeImg = false
	}

	if !validTypeImg {
		return &entity.ApplicationError{
			Err:        []error{errors.New("cannot upload image type. must PNG/JPEG/JPG")},
			HTTPStatus: http.StatusBadRequest,
		}
	}

	for _, key := range users {
		api := anaconda.NewTwitterApiWithCredentials(key.AccessToken, key.AccessTokenSecret, os.Getenv("TWITTER_CONSUMER_KEY"), os.Getenv("TWITTER_CONSUMER_SECRET"))
		media, err := api.UploadMedia(colImg[1])
		if err != nil {
			return &entity.ApplicationError{
				Err:        []error{errors.New("something when wrong. can't upload media on twitter")},
				HTTPStatus: http.StatusBadRequest,
			}
		}

		v := url.Values{}
		v.Set("media_ids", strconv.FormatInt(media.MediaID, 10))

		tweet, err := api.PostTweet(request.Text, v)
		if err != nil {
			return &entity.ApplicationError{
				Err:        []error{errors.New("something when wrong. can't post tweet")},
				HTTPStatus: http.StatusBadRequest,
			}
		}

		requestCreate := entity.OFeedTwitter{
			UserID:      request.UserID,
			Text:        request.Text,
			UserTweetID: key.UserTweetID,
		}

		if err := providerKeyword.CreateTweetPostOauth(ctx, requestCreate, tweet.IdStr, tweet.User.ScreenName); err != nil {
			return err
		}
	}

	return nil
}
