package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
    //数据库
    "database/sql"
    "fmt"
    _"github.com/go-sql-driver/mysql"

)

//var userIdSequence = int64(1)

func Register(c *gin.Context) {

	username := c.Query("username")
	password := c.Query("password")

	//数据库验证是否存在该username
    //dsn := "user:password@tcp(127.0.0.1:3306)/douyin"
    db, err := sql.Open("mysql", dsn)
    if err != nil {
		panic(err)
	}

    sqlStr := "select password from users where name = ?"
    var Isexist string
     err= db.QueryRow(sqlStr, username).Scan(&Isexist)

    if err == nil {
       c.JSON(http.StatusOK, UserLoginResponse{
         Response: Response{StatusCode: 1 ,
         StatusMsg: "User already exist",
         },   
       })
    }else{
	   // atomic.AddInt64(&userIdSequence, 1) //自增函数
       // NewUserID:=int64(userIdSequence)
       // string := strconv.FormatInt(userIdSequence,10)
        //TokenUserid,err := strconv.Atoi(string)
        sqlStr := "insert into users(name,password,follow_count,follower_count,total_favorited,work_count,favorite_count) values (?,?,?,?,?,?,?)"
        _,err = db.Exec(sqlStr,username,password,0,0,0,0,0)
        
        if err != nil {
		    fmt.Printf("insert failed, err:%v\n", err)
	    }
        var NewUserID int
        err=db.QueryRow("select User_id from users order by User_id desc limit 1").Scan(&NewUserID)
        //userIdSequence = int64(userIdSequence)
        NewUserID64:=int64(NewUserID)
	    c.JSON(http.StatusOK, UserLoginResponse{
		    Response: Response{
                StatusCode: 0,
            StatusMsg: "Successfully",
        },
		    //UserId:   NewUserID64,
		    Token:    SettingToken(NewUserID),
            UserId:   NewUserID64,
	    })
    }
}
