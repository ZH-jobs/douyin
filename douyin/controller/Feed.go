package controller

import (
    "github.com/gin-gonic/gin"
	"net/http"
	"time"
    //数据库
    "database/sql"
    _"github.com/go-sql-driver/mysql"
    "fmt"
)

type FeedResponse struct {
	Response
	VideoList []VideoReturn `json:"video_list,omitempty"`
	NextTime  int64   `json:"next_time,omitempty"`
}

/*type MysqlVideo struct{
    Video_id        int64
    User_id         int64
    play_url        string
    cover_url       string
    title           string
    favourite_count int64
    comment_count   int64
    date            string
}*/

// Feed same demo video list for every request
func Feed(c *gin.Context) {
    Current_Token:=c.Query("token")
    var user_id int =IsTrueToken(Current_Token)
    if user_id == -1{//Token无法解析，需要重新注册和登陆
        c.JSON(http.StatusOK,FeedResponse{
            Response: Response{StatusCode: 1,StatusMsg: "Token is wrong"},
        })
    }else{
        
        //dsn := "user:password@tcp(127.0.0.1:3306)/douyin"//数据库基本信息
        db, err := sql.Open("mysql", dsn)
        if err != nil {
            panic(err)
        }
        //数据库查询后30条视频数据
        rows,err:=db.Query("select Video_id,User_id,play_url,cover_url,title,favorite_count,comment_count from video order by Video_id desc limit 30")
        if err!=nil{
            c.JSON(http.StatusOK,FeedResponse{
                Response: Response{StatusCode: 1,StatusMsg: "database is empty"},
            })
        }else{
            var Videos  []VideoReturn//结构体数组
            defer rows.Close()
            for rows.Next() {
                //视频基本信息定义
                var Video_id int64=0
                var User_id  int64=0
                var play_url string
                var cover_url string
                var title     string
                var favorite_count  int64
                var comment_count  int64
                /////////////////
		        err := rows.Scan(&Video_id,&User_id,&play_url,&cover_url,&title,&favorite_count,&comment_count)//查询mysql中video
		        if err != nil {
			        fmt.Printf("scan failed, err:%v\n", err)
			        return
		        }
                //用户是否关注该作者
                sqlStr := "select * from follow_list  where User_id=? and follow_id= ? "
                var istruefollow bool =true
                err=db.QueryRow(sqlStr,user_id,User_id).Scan(&istruefollow) 
                if err!=nil{
                    istruefollow =false
                }
                //用户是否喜欢该视频
                var istruefavourite bool =true
                sqlStr = "select * from favourite_list  where User_id=? and Video_id= ?"
                err=db.QueryRow(sqlStr,user_id,Video_id).Scan(&istruefavourite)
                if err!=nil{
                    istruefavourite =false
                }
                //作者关注数量与粉丝数量
                sqlStr = "select name,follow_count,follower_count,work_count,favorite_count from users where User_id=?"
                var followcount,followercount,work_count,user_favorite_count int64=0,0,0,0
                var name string
                err=db.QueryRow(sqlStr,User_id).Scan(&name,&followcount,&followercount,&work_count,&user_favorite_count)
                //作者姓名
                //需要返回的User结构体
                var tempUser =UserReturn{
                    FollowCount:     followcount,
                    FollowerCount:   followercount,
                    ID:              User_id,
                    IsFollow:        istruefollow,
                    Name:            name,
                    WorkCount:       work_count,
                    FavoriteCount:  user_favorite_count,
                }
                //以下定义Video的临时结构体
                var tempVideoMax =VideoReturn{
                    Author:          tempUser,
                    CommentCount:    comment_count,
                    CoverURL:        cover_url,
                    FavoriteCount:   favorite_count,
                    ID:              Video_id,
                    IsFavorite:      istruefavourite,
                    PlayURL:         play_url,
                    Title:           title,
                }
                //将临时Video结构体放入结构体数组中
               Videos =  append(Videos,tempVideoMax)
	        } 
    	    c.JSON(http.StatusOK, FeedResponse{
		        Response:  Response{StatusCode: 0},
		        VideoList: Videos,
		        NextTime:  time.Now().Unix(),
	        })
        }
    }
}
