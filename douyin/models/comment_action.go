package models

// CommentAction 评论表
type CommentAction struct {
	Id         int64  `json:"id,omitempty"`
	CommentId  int64  `json:"comment_id,omitempty"`
	VideoId    int64  `json:"video_id,omitempty"`
	UserId     int64  `json:"user_id,omitempty"`
	Content    string `json:"content,omitempty"`
	CreateTime string `gorm:"column:comment_time" json:"creat_time,omitempty"`
}

func (ct CommentAction) TableName() string {
	return "comment_action"
}
