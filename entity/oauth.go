package entity

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

type OUserTwitterInfo struct {
	UserTweetID     string `json:"user_tweet_id"`
	Name            string `json:"name"`
	Username        string `json:"username"`
	ProfileImageURL string `json:"profile_image_url"`
}
