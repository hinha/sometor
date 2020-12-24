package entity

type TwitterResult struct {
	Results    []TwitterItem `json:"results"`
	LastUpdate string        `json:"last_update"`
}

type TwitterItem struct {
	Id                  int64   `json:"id"`
	Permalink           string  `json:"permalink"`
	Timestamp           int64   `json:"timestamp"`
	CreatedAt           string  `json:"created_at"`
	CreatedAtTime       string  `json:"created_at_time"`
	FullTextNorm        string  `json:"full_text_norm"`
	FullTextClean       string  `json:"full_text_clean"`
	FullText            string  `json:"full_text"`
	TextSentiment       string  `json:"text_sentiment"`
	UserId              string  `json:"user_id"`
	UserName            string  `json:"user_name"`
	ScreenName          string  `json:"screen_name"`
	UserDescription     string  `json:"user_description"`
	UserVerified        bool    `json:"user_verified"`
	UserCreatedAt       string  `json:"user_created_at"`
	UserFollowersCount  int32   `json:"user_followers_count"`
	UserFriendsCount    int32   `json:"user_friends_count"`
	UserStatusesCount   int32   `json:"user_statuses_count"`
	UserFavouritesCount int32   `json:"user_favourites_count"`
	UserListedCount     int32   `json:"user_listed_count"`
	UserMediaCount      int32   `json:"user_media_count"`
	UserLocation        string  `json:"user_location"`
	UserProtected       bool    `json:"user_protected"`
	UserLinkUrl         string  `json:"user_link_url"`
	UserProfileImageUrl string  `json:"user_profile_image_url"`
	ReplyCount          int32   `json:"reply_count"`
	RetweetCount        int32   `json:"retweet_count"`
	LikeCount           int32   `json:"like_count"`
	QuoteCount          int32   `json:"quote_count"`
	ConversationId      int64   `json:"conversation_id"`
	Device              string  `json:"device"`
	Lang                string  `json:"lang"`
	Engagement          int64   `json:"engagement"`
	PotentialReach      int64   `json:"potential_reach"`
	EngagementRate      float32 `json:"engagement_rate"`
	PlaceType           string  `json:"place_type"`
	PlaceName           string  `json:"place_name"`
	CountryCode         string  `json:"country_code"`
	BoundingBoxType     string  `json:"bounding_box_type"`
	Media               []struct {
		Type   string          `json:"type"`
		Source []TwitterSource `json:"source"`
	} `json:"media"`
	BoundingBoxCoordinates []struct {
		Longitude float32 `json:"longitude"`
		Latitude  float32 `json:"latitude"`
	} `json:"bounding_box_coordinates"`
	Mentions []struct {
		Id          int64  `json:"id"`
		Text        string `json:"text"`
		DisplayName string `json:"display_name"`
	} `json:"mentions"`
	Hashtags []struct {
		Text string `json:"text"`
	} `json:"hashtags"`
}

type TwitterSource struct {
	PreviewUrl   string `json:"preview_url"`
	FullUrl      string `json:"full_url"`
	ThumbnailUrl string `json:"thumbnail_url"`
	Duration     string `json:"duration"`
}
