package controller

import "time"
type Response struct {
     StatusCode     int64  `json:"status_code"`
     StatusMsg      string `json:"status_msg"`
}
//用于后端返回User结构体
type UserReturn struct {
     FollowCount    int64  `json:"follow_count"`
     FollowerCount  int64  `json:"follower_count"`
     ID             int64  `json:"id"`
     IsFollow       bool   `json:"is_follow"`
     Name           string `json:"name"`
     WorkCount      int64  `json:"work_count"`
     FavoriteCount  int64  `json:"favourite_count"`

}
//
type User struct {
	FavoriteCount int64  `json:"favorite_count,omitempty"`
	//IsFavorite    bool   `json:"is_favorite,omitempty"`
	CommentCount  int64  `json:"comment_count,omitempty"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
	ID            int64    `json:"id"`
	IsFollow      bool   `json:"is_follow"`
	Name          string `json:"name"`
}
//用于后端返回Video结构体
type VideoReturn struct {
     Author         UserReturn   `json:"author"`
     CommentCount   int64  `json:"comment_count"`
     CoverURL       string `json:"cover_url"`
     FavoriteCount  int64  `json:"favorite_count"`
     ID             int64  `json:"id"`
     IsFavorite     bool   `json:"is_favorite"`
     PlayURL        string `json:"play_url"`
     Title          string `json:"title"`
     Date           time.Time `json:"date"`
}

// Video 结构体表示一个抖音视频
type Video struct {
	Video_id       int64     `gorm:"primarykey;AUTO_INCREMENT"`
	User_id        int64     `gorm:"column:User_id"`
	Play_url       string    `gorm:"column:play_url"`
	Cover_url      string    `gorm:"column:cover_Url"`
	Title          string    `gorm:"column:title"`
	Favorite_count int64     `gorm:"column:favorite_count"`
	Comment_count  int64     `gorm:"column:comment_count"`
	Date           time.Time `gorm:"column:date"`
	//gorm.Model               // gorm 模型，包含 ID、CreatedAt、UpdatedAt、DeletedAt 等字段
}

type VideoReturnNew struct {
     Author         UserReturn   `json:"author"`
     CommentCount   int64  `json:"comment_count"`
     CoverURL       string `json:"cover_url"`
     FavoriteCount  int64  `json:"favorite_count"`
     ID             int64  `json:"id"`
     IsFavorite     bool   `json:"is_favorite"`
     PlayURL        string `json:"play_url"`
     Title          string `json:"title"`
}
/*type Message struct {
     Content        string `json:"content"`
     CreateTime     time `json:"create_time"`
     ID             int64  `json:"id"`
}*/
// Message 结构体表示一条聊天记录
type Message struct {
	Message_id   int64 `json:"Message_id"`
	From_user_id int    `json:"from_user_id"`
	To_user_id   int    `json:"to_user_id"`
	Content      string  `json:"content"`
	Creat_time   time.Time  `json:"creat_time"`
}
/*type Comment struct {
     Content        string  `json:"content"`
     CreateDate     string  `json:"create_date"`
     ID             int64   `json:"id"`
     User           User    `json:"user"`
}*/
