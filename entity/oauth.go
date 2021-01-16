package entity

// LoginWToken got request from frontend
type LoginWToken struct {
	TokenID string `json:"tokenId"`
}

type OUserTwitter struct {
	Name              string `json:"name"`
	Username          string `json:"username"`
	ProfileImageURL   string `json:"profile_image_url"`
	UserID            string `json:"user_id"`
	AccessToken       string `json:"access_token"`
	AccessTokenSecret string `json:"access_token_secret"`
	ClientID          string `json:"client_id"`
}

type OUserTwitterInfo struct {
	UserID          string `json:"user_id"`
	Name            string `json:"name"`
	Username        string `json:"username"`
	ProfileImageURL string `json:"profile_image_url"`
}
