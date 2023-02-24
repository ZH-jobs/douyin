package controller

import (
	"fmt"
	"net/http"
	"path/filepath"
	"time"
     "log"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
    "database/sql"
    _"github.com/go-sql-driver/mysql"
    "strconv"
)

// VideoList 包含了多个视频的数组
type VideoList1 struct {
	Response
	Videos []Video `json:"videos"`
}

// 初始化数据库连接
func setupDB() (*gorm.DB, error) {

	// 连接 MySQL 数据库，格式为 "username:password@tcp(host:port)/database_name?charset=utf8mb4&parseTime=True&loc=Local"
	//dsn := "root:password@tcp(127.0.0.1:3306)/video_db?charset=utf8mb4&parseTime=True&loc=Local"
	return gorm.Open(mysql.Open(dsn))
}

func (video *Video) TableName() string {
	return "video" //你希望他的名字是什么就填什么
}

// 定义Publish函数，用于处理视频发布请求
func Publish(c *gin.Context) {
	// 从请求中获取token
	token := c.PostForm("token")

	// 解析Token得到用户id
	userID := IsTrueToken(token)
    log.Println("123:",token)
	if userID == -1 {
		// 返回状态码为1，表示用户不存在
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

	// 从请求中获取文件
	data, err := c.FormFile("data")

	if err != nil {
		// 返回错误信息
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	// 获取文件名
	filename := filepath.Base(data.Filename)

	// 构造最终文件名
	finalName := fmt.Sprintf("%d_%s", userID, filename)

	// 构造文件存储路径
	saveFile := filepath.Join("./public/", finalName)

	// 保存文件
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		// 返回错误信息
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	// 返回状态码为0，表示文件上传成功
	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  finalName + " uploaded successfully",
	})

	//获取封面名
	cover, err := c.FormFile("cover")
	if err != nil {
		// 返回错误信息
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	// 获取封面名
	covername := filepath.Base(cover.Filename)

	// 构造最终封面名
	finalcoverName := fmt.Sprintf("%d_%s", userID, covername)

	// 构造文件存储路径
	savecoverFile := filepath.Join("./cover/", finalcoverName)

	// 保存封面
	if err := c.SaveUploadedFile(data, savecoverFile); err != nil {
		// 返回错误信息
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	// 返回状态码为0，表示封面上传成功
	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  finalcoverName + " uploaded successfully",
	})

	// 保存视频信息到数据库
	db, err := setupDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
    
	// 自动创建Video结构的表
	db.AutoMigrate(&Video{})
    userIDnew:=int64(userID)
	video := Video{
		User_id: userIDnew,
		//Author:        User{Name: user.Name},
		Title:     finalName,
		Date:      time.Now(),
		Play_url:  "http://localhost:8080/douyin/public/" + finalName,
		Cover_url: "http://localhost:8080/douyin/cover/" + finalcoverName,
	}
    log.Println("Video :",video)

	if err := db.Table("video").Create(&video).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Video uploaded successfully!"})
}



// GetVideos 返回视频列表
func PublishList(c *gin.Context) {
	// 使用 Gorm 和 MySQL 连接数据库
	/*db, err := setupDB()
	if err != nil {
		fmt.Println("Failed to connect to database:", err)
		//return VideoList{}
	}

	// 查询视频列表
	var videos []Video
	db.Find(&videos)*/
    /*
    userid:=c.Query("user_id")
    userID,err := strconv.ParseInt(userid, 10, 64)
    db, err := sql.Open("mysql", dsn)
    if err != nil {
		panic(err)
	}
    sqlStr := "select * from video where user_id = ?"
    rows,err:=db.Query(sqlStr,userID)
    for rows.Next(){
        
    }
    */
    userid:=c.Query("user_id")
    userID,err:= strconv.ParseInt(userid, 10, 64)
    if err!=nil{
       fmt.Printf("1234567")
    }
    Current_Token:=c.Query("token")
    var user_id int =IsTrueToken(Current_Token)
    if user_id == -1{//Token无法解析，需要重新注册和登陆
        c.JSON(http.StatusOK,FeedResponse{
            Response: Response{StatusCode: 1,StatusMsg: "Token is wrong"},
        })
    }else{
        db, err := sql.Open("mysql", dsn)
        if err != nil {
            panic(err)
        }
        //log.Println("12345:",userID)
        sqlStr:="select Video_id,User_id,play_url,cover_url,title,favorite_count,comment_count from video where user_id = ?"
        rows,err:=db.Query(sqlStr,userID)

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
                sqlStr = "select * from follow_list  where User_id=? and follow_id= ? "
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
                Videos =  append(Videos,tempVideoMax)
            }
                c.JSON(http.StatusOK, FeedResponse{
		        Response:  Response{StatusCode: 0},
		        VideoList: Videos,
	        })
        }    
    }
	/*c.JSON(http.StatusOK, VideoList1{
		Response: Response{
			StatusCode: 0,
			StatusMsg:  "获取视频列表成功",
		},
		Videos: videos,
	})*/
}

