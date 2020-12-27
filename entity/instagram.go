package entity

type InstagramResult struct {
	Results    []InstagramItem `json:"results"`
	LastUpdate string          `json:"last_update"`
}

type InstagramItem struct {
	Id                  string `json:"id"`
	Permalink           string `json:"permalink"`
	Timestamp           int64  `json:"timestamp"`
	CreatedAt           string `json:"created_at"`
	CreatedAtTime       string `json:"created_at_time"`
	TextClean           string `json:"text_clean"`
	Text                string `json:"text"`
	TextSentiment       string `json:"text_sentiment"`
	UserId              string `json:"user_id"`
	UserName            string `json:"user_name"`
	Name                string `json:"name"`
	UserDescription     string `json:"user_description"`
	UserVerified        bool   `json:"user_verified"`
	UserFollowersCount  int64  `json:"user_followers_count"`
	UserFollowingCount  int64  `json:"user_following_count"`
	UserTimelineCount   int32  `json:"user_timeline_count"`
	UserHighlightCount  int32  `json:"user_highlight_count"`
	IsPrivate           bool   `json:"is_private"`
	UserLinkUrl         string `json:"user_link_url"`
	UserProfileImageUrl string `json:"user_profile_image_url"`
	LikeCount           int64  `json:"like_count"`
	CommentCount        int64  `json:"comment_count"`
	VideoViewCount      int64  `json:"video_view_count"`
	Mentions            []struct {
		Text string `json:"text"`
	} `json:"mentions"`
	Hashtags []struct {
		Text string `json:"text"`
	} `json:"hashtags"`
	Engagement     int64   `json:"engagement"`
	PotentialReach int64   `json:"potential_reach"`
	EngagementRate float32 `json:"engagement_rate"`
	CommentDisable bool    `json:"comment_disable"`
	IsVideo        bool    `json:"is_video"`
	VideoUrl       string  `json:"video_url"`
	DisplayUrl     string  `json:"display_url"`
}
