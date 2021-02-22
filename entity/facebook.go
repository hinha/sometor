package entity

type FacebookResult struct {
	Results    []FacebookItem `json:"results"`
	LastUpdate string         `json:"last_update"`
}

type FacebookItem struct {
	Comments        int    `json:"comments"`
	CreatedAt       string `json:"created_at"`
	Image           string `json:"image"`
	Likes           string `json:"likes"`
	Shares          int    `json:"shares"`
	Permalink       string `json:"permalink"`
	PostID          string `json:"post_id"`
	SharedLink      string `json:"shared_link"`
	SharedText      string `json:"shared_text"`
	StrUpdatedDate  string `json:"str_updated_date"`
	Text            string `json:"text"`
	TextSentiment   string `json:"text_sentiment"`
	Timestamp       int64  `json:"timestamp"`
	UserDisplayName string `json:"user_display_name"`
	UserName        string `json:"user_name"`
	UserProfilePic  string `json:"user_profile_pic"`
}
