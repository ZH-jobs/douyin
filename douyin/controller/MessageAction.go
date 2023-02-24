package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
    "database/sql"
    _"github.com/go-sql-driver/mysql"
    "strconv"
)

// MessageAction no practical effect, just check if token is valid
func MessageAction(c *gin.Context) {
	token := c.Query("token")
	toUserIdstring := c.Query("to_user_id")//string类型，需要转化为int64
	content := c.Query("content")
	//数据库检测Token是否存在
    user_ID:=IsTrueToken(token)
    //dsn :="user:password@tcp(127.0.0.1:3306)/douyin"
    db,err:=sql.Open("mysql",dsn)
    if err !=nil{
        panic(err)
    }
    sqlStr:="select name from users where User_id =? "
    toUserIdint64, err := strconv.ParseInt(toUserIdstring, 10, 64)
    var isexisttemp1  string//发送方是否存在
    var isexist  string//检测发送目标用户是否存在
    err1:=db.QueryRow(sqlStr,user_ID).Scan(&isexisttemp1)//检测用户是否存在
    err2:=db.QueryRow(sqlStr,toUserIdint64).Scan(&isexist)//检测发送目标是否存在

    if err1!=nil||err2!=nil{
        c.JSON(http.StatusOK, Response{StatusCode: 0, StatusMsg: "User doesn't exist"}) 
    }else{
        //插入聊天信息到数据库中
        sqlStr := "insert into message_list(Message_id, to_user_id,from_user_id,content,creat_time) values (?,?,?,?,?)"
        _,err:= db.Exec(sqlStr,0,toUserIdint64,user_ID,content,time.Now().Format("2006-01-02"))
        if err == nil{
            c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "That's OK"})
        }else{
            c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "Insert failed"})
        }
        defer db.Close()
    }
}

