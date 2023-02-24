package controller

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// GetDB 获取数据库连接
func GetDB() (*gorm.DB, error) {
    
	//url := "root:root@tcp(127.0.0.1:3306)/douyin?charset=utf8mb4&parseTime=True&loc=Local"
    url:=dsn+"?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(url), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			//使用单数列表名
			// 不加这段代码 框架会在你指定的列表名后面添加 s 导致出错
			SingularTable: true,
		},
	})

	return db, err
}
