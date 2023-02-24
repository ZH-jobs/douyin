package models

type Users struct {
	UserId   string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Password string `json:"password,omitempty"`
	//喜欢视频总数
	FavoriteCount int   `json:"favourit_count"`
	FollowerCount int64 `json:"follower_count"`
	FollowCount   int64 `json:"follow_count"`
	//他人喜欢总数
	TotalFavorited int64 `json:"total_favorited"`
	WorkCount      int64 `json:"work_count"`
}
