package main

import "fmt"

//// User2 一对一关系一般用于表的扩展,拆分为两张表，常用的字段放主表，不常用的字段放详情表
//type User2 struct {
//	ID       uint
//	Name     string
//	Age      int
//	Gender   bool
//	UserInfo *UserInfo // 通过UserInfo可以拿到用户详情信息
//}
//
//type UserInfo struct {
//	User2ID uint // 外键
//	User2   User2
//	ID      uint
//	Addr    string
//	Like    string
//}

type User3 struct {
	ID        uint
	Name      string
	Age       int
	Gender    bool
	UserInfo3 UserInfo3 // 通过UserInfo可以拿到用户详情信息
}

type UserInfo3 struct {
	User3   *User3 // 要改成指针，不然就嵌套引用了
	User3ID uint   // 外键
	ID      uint
	Addr    string
	Like    string
}

func main() {
	//DB.Debug().AutoMigrate(&User2{}, &UserInfo{})

	//// 添加用户，自动添加用户详情
	//// 添加用户详情，关联已有用户
	//DB.Create(&User2{
	//	Name:   "枫枫",
	//	Age:    21,
	//	Gender: true,
	//	UserInfo: &UserInfo{
	//		Addr: "湖南省",
	//		Like: "写代码",
	//	},
	//})
	//// 添加附表
	//DB.Create(&UserInfo{
	//	User2ID: 2,
	//	Addr:    "南京市",
	//	Like:    "吃饭",
	//})

	//DB.Debug().AutoMigrate(&UserInfo3{}, &User3{})
	//DB.Create(&User3{
	//	Name:   "枫枫",
	//	Age:    21,
	//	Gender: true,
	//	UserInfo3: UserInfo3{
	//		Addr: "湖南省",
	//		Like: "写代码",
	//	},
	//})
	//// 添加附表
	//DB.Create(&UserInfo3{
	//	User3ID: 2,
	//	Addr:    "南京市",
	//	Like:    "吃饭",
	//})

	// 不限于重新迁移，直接添加即可
	var user User3
	//DB.Take(&user, 2)
	//DB.Create(&UserInfo3{
	//	User3: &user,
	//	Addr:  "南京市",
	//	Like:  "吃饭",
	//})
	//
	// 查询
	// 通过主表查副表
	DB.Preload("UserInfo3").Take(&user)
	fmt.Println(user)
}
