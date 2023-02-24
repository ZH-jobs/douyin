/* ************************************************************************
> File Name:     Login.go
> Author:        周昊
> Created Time:  2023年02月18日 星期六 14时33分19秒
> Description:   用户登陆
 ************************************************************************/

package controller

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "database/sql"
    _"github.com/go-sql-driver/mysql"
)

type UserLoginResponse struct {
   Response 
   UserId int64  `json:"user_id,omitempty"`
   Token  string `json:"token"`
}
//登陆函数
func Login(c *gin.Context){

    username:=c.Query("username")
    password:=c.Query("password")
    //连接数据库
    //dsn:="user:password@tcp(127.0.0.1:3306)/douyin"
    db,err:= sql.Open("mysql",dsn)
    if err!=nil{
        panic(err)
    }
    //根据账号密码查询数据库
    sqlStr :="select User_id from users where name = ? and password = ?"
    var User_id int
    err =db.QueryRow(sqlStr,username,password).Scan(&User_id)
    if err!=nil{
        //未找到用户
        c.JSON(
            http.StatusOK,
            UserLoginResponse{
                Response:Response{
                    StatusCode: 1,
                    StatusMsg: "User doesn't exist",
                },
            },
        )
    }else{
        //找到用户返回Json
        Token:=SettingToken(User_id)
        User_id64:=int64(User_id)
        c.JSON(
            http.StatusOK,
            UserLoginResponse{
                Response:   Response{StatusCode: 0},
                UserId:     User_id64,
                Token:      Token,
        })
    }
}
