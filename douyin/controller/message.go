package controller

import (
	"net/http"
	"strconv"
    "log"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 连接MySQL数据库
func dbConnect() (*gorm.DB, error) {
	//dsn := "root:password@tcp(localhost:3306)/chat?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		return nil, err
	}
	return db, nil
}

// 表名为Message_list
func (message *Message) TableName() string {
	return "message_list"
}

// 声明一个 ChatResponse 结构体，用于作为返回值返回聊天记录和请求状态。
type ChatResponse struct {
	Response
	MessageList []Message `json:"message_list"`
}

// MessageAction 处理客户端发送的新聊天消息的请求
/*func MessageAction(c *gin.Context) {
	// 获取请求中的参数
	token := c.PostForm("token")         // 获取请求中的 token 参数，用于用户身份认证
	toUserId := c.PostForm("to_user_id") // 获取请求中的 to_user_id 参数，用于指定聊天对象
	content := c.PostForm("content")     // 获取请求中的 content 参数，用于聊天消息的内容

	// 解析token，返回用户id
	userID := IsTrueToken(token)
	if userID == -1 {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"}) // 返回错误状态码和信息。
	}
	userIdB, _ := strconv.Atoi(toUserId) // 将 to_user_id 参数转换成整数类型

	//连接数据库
	db, err := dbConnect()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//自动创建一个Message表
	db.AutoMigrate(&Message{})

	msgRecord := Message{
		From_user_id: userID,
		To_user_id:   int(userIdB),
		Content:      content,
		Creat_time:   time.Now(),
	}

	// 保存消息记录到数据库中
	result := db.Create(&msgRecord)
	if result.Error != nil {
		return
	}

	c.JSON(http.StatusOK, Response{StatusCode: 0}) // 返回成功状态码。
}*/

// MessageChat 处理客户端获取聊天消息列表的请求
func MessageChat(c *gin.Context) {
	// 获取请求中的参数
	token := c.Query("token")
	toUserId := c.Query("to_user_id")
	// 解析token，返回用户id
	userID := IsTrueToken(token)
    log.Println("用户id",userID)
	if userID != -1 {
		userIdB, _ := strconv.Atoi(toUserId)

		db, err := dbConnect()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
        log.Println("用户B：",toUserId)
		var messages []Message
        var messages2 []Message
		db.Omit("id").Where("from_user_id = ? ", userID).Where("to_user_id = ?", userIdB).Find(&messages)
        db.Omit("id").Where("from_user_id = ? ", userIdB).Where("to_user_id = ?", userID).Find(&messages2)
        for i:=0;i<len(messages2);i++ {
            messages=append(messages,messages2[i])
        }
        //messages=append(messages,messages2)
        //db.Omit("id").Where("from_user_id =  ", userID).Where("to_user_id = ?", userIdB).Find(&messages)
		// 返回一个包含聊天记录的 ChatResponse 结构体，其中 StatusCode 为 0 表示成功，MessageList 是一个切片，包含了当前聊天对象的所有聊天记录。
		c.JSON(http.StatusOK, ChatResponse{Response: Response{StatusCode: 0}, MessageList: messages}) // 返回成功状态码和聊天记录中的消息列表
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"}) // 返回失败状态码和错误信息
	}
}
