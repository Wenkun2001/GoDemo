package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB
var mysqlLogger logger.Interface

func init() {
	// MySQL 配置信息
	username := "root"   // 账号
	password := "010226" // 密码
	host := "127.0.0.1"  // 地址
	port := 3306         // 端口
	DBname := "gorm"     // 数据库名称
	timeout := "10s"     // 连接超时，10秒

	mysqlLogger = logger.Default.LogMode(logger.Info)

	//// 自定义日志的显示
	//mysqlLogger := logger.New(
	//	// （日志输出的目标，前缀和日志包含的内容
	//	log.New(os.Stdout, "\r\n", log.LstdFlags),
	//	logger.Config{
	//		SlowThreshold:             time.Second, // 慢SQL阈值
	//		LogLevel:                  logger.Info, // 日志级别
	//		IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNoutFound（记录未找到）错误
	//		Colorful:                  true,        // 使用彩色打印
	//	},
	//)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", username, password, host, port, DBname, timeout)
	// Open 连接
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			//TablePrefix:   "f_",  //表名前缀
			SingularTable: true,  // 单数表名
			NoLowerCase:   false, // 是否关闭小写转换
		},
		//Logger: mysqlLogger, //全局日志
	})
	if err != nil {
		panic("failed to connect mysql." + err.Error())
	}
	DB = db
}

//type Student struct {
//	// 模型定义
//	ID      uint    `gorm:"size:10"` // 默认作为主键
//	Name    string  `gorm:"type:varchar(16)"`
//	Age     int     `gorm:"size:3"`
//	Address *string `gorm:"size:128"` //使用指针可以存储空值
//	Type    string  `gorm:"column:_type;size:4"`
//	Date    string  `gorm:"default:2023;comment:日期"`
//}

//func main() {
//	//DB = DB.Session(&gorm.Session{
//	//	Logger: mysqlLogger,
//	//})
//
//	// 自动生成表结构
//	//DB.AutoMigrate(&Student{})
//	带日志生成表结构
//	DB.Debug().AutoMigrate(&Student{})
//}
