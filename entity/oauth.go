package entity

import "time"

// LoginWToken got request from frontend
type LoginWToken struct {
	TokenID string `json:"tokenId"`
}

type OUserTwitter struct {
	Name              string `validate:"min=1,max=100" json:"name"`
	Username          string `validate:"min=1,max=100" json:"username"`
	ProfileImageURL   string `validate:"min=1,max=255" json:"profile_image_url"`
	UserID            string `validate:"min=1,max=100" json:"user_id"`
	AccessToken       string `validate:"min=1,max=150" json:"access_token"`
	AccessTokenSecret string `validate:"min=1,max=150" json:"access_token_secret"`
	ClientID          string `validate:"min=1,max=100" json:"client_id"`
}

type OUserTwitterKey struct {
	UserTweetID       string `json:"user_tweet_id"`
	AccessToken       string `json:"access_token"`
	AccessTokenSecret string `json:"access_token_secret"`
}

type OUserTwitterInfo struct {
	UserTweetID     string `json:"user_tweet_id"`
	Name            string `json:"name"`
	Username        string `json:"username"`
	ProfileImageURL string `json:"profile_image_url"`
}

type OFeedTwitter struct {
	UserID      string `json:"user_id"`
	UserTweetID string `json:"user_tweet_id"`
	Text        string `json:"text"`
}

// type multi
type OFeedTwitterAll struct {
	UserID string `json:"user_id"`
	Text   string `json:"text"`
}

type OFeedTwitterInfo struct {
	StatusesID string    `json:"statuses_id"`
	Text       string    `json:"text"`
	Permalink  string    `json:"permalink"`
	CreatedAt  time.Time `json:"created_at"`
}
