package main

//
//import (
//	"fmt"
//	"gorm.io/gorm"
//)
//
//type Student2 struct {
//	ID      uint    `gorm:"size:10"` // 默认作为主键
//	Name    string  `gorm:"type:varchar(16)"`
//	Age     int     `gorm:"size:3"`
//	Address *string `gorm:"size:128"` //使用指针可以存储空值
//	Gender  bool    `gorm:"default:true"`
//}
//
//// 创建HOOK钩子函数
//func (user *Student2) BeforeCreate(tx *gorm.DB) (err error) {
//	address := fmt.Sprintf("%s@qq.com", user.Name)
//	user.Address = &address
//	return nil
//}
//
//func main() {
//	//DB.AutoMigrate(&Student2{})
//
//	//email := "12345"
//	//// insert 添加记录
//	//s1 := Student2{
//	//	// ID自动生成
//	//	Name:    "lwk",
//	//	Age:     21,
//	//	Gender:  true,
//	//	Address: &email,
//	//}
//
//	//// 批量插入
//	//var StudentList []Student2
//	//for i := 0; i < 10; i++ {
//	//	StudentList = append(StudentList, Student2{
//	//		// ID自动生成
//	//		Name:    fmt.Sprintf("lwk%d", i+1),
//	//		Age:     21 + i + 1,
//	//		Gender:  true,
//	//		Address: &email,
//	//	})
//	//}
//	////err := DB.Create(&s1).Error
//	//err := DB.Create(&StudentList).Error
//	//log.Println(err)
//
//	// 单条记录查询
//	//var student Student2
//	DB = DB.Session(&gorm.Session{
//		Logger: mysqlLogger,
//	})
//	//DB.Take(&student)
//	//fmt.Println(student)
//	//DB.First(&student)
//	//fmt.Println(student)
//	//DB.Last(&student)
//	//fmt.Println(student)
//
//	//// Take默认根据主键查询，可以是字符串，可以是数字
//	//err := DB.Take(&student, "45").Error
//	//fmt.Println(err)
//	//fmt.Println(student)
//
//	//// 使用？作为占位符， 将查询的内容放入？
//	//err := DB.Take(&student, "name = ?", "lwk").Error
//	//fmt.Println(err)
//	//fmt.Println(student)
//
//	//// 将参数全部转义，#前;或者空格，每个参数’隔开后再用or等关系
//	//err := DB.Take(&student, fmt.Sprintf("name = '%s'", "lwk' or 1=1;#")).Error
//	////err := DB.Take(&student, fmt.Sprintf("name = '%s'", "lwk' or 1=1 #")).Error
//	//fmt.Println(err)
//	//fmt.Println(student)
//
//	//// 获取查询结果
//	//// 获取查询记录数
//	//count := DB.Find(&student).RowsAffected
//	//fmt.Println(count)
//	//
//	//// 是否查询失败（查询为空， 查询条件错误， sql语法错误）
//	//err := DB.Take(&student, "xx").Error
//	//switch err {
//	//case gorm.ErrRecordNotFound:
//	//	fmt.Println("没有找到")
//	//default:
//	//	fmt.Println("sql错误")
//	//}
//
//	//// 查询多条记录
//	//var studentList []Student2
//	//DB.Find(&studentList)
//	//for _, student := range studentList {
//	//	fmt.Println(student)
//	//	// 指针类型看不到实际的内容
//	//	// 序列化之后，会转换为我们可以看得懂的方式
//	//	data, _ := json.Marshal(student)
//	//	fmt.Println(string(data))
//	//}
//	//
//	//// 根据主键列表查询
//	//DB.Find(&studentList, []int{4, 7, 9})
//	//fmt.Println(studentList)
//	//
//	//// 根据其他条件查询
//	//DB.Find(&studentList, "name in ?", []string{"lwk", "lwk2"})
//	//fmt.Println(studentList)
//
//	//// 更新
//	//// 更新的前提是先查询到记录
//	//var student Student2
//	//DB.Take(&student)
//	//student.Age = 23
//	//// 全字段更新, 零值、nil值也会更新
//	//DB.Save(&student)
//	//
//	//// 批量更新
//	//var studentList []Student2
//	//// 方法一
//	//DB.Find(&studentList, "age = ?", 21).Update("email", "is")
//	//// 方法二
//	//DB.Model(&Student2{}).Where("age = ?", 22).Update("email", "is")
//	//
//	//// 更新多列
//	//address := "xx@qq.com"
//	//// 如果是结构体，默认不更新零值
//	//DB.Model(&Student2{}).Where("age = ?", 21).Updates(Student2{
//	//	Address: address,
//	//	Gender:  false,
//	//})
//	//
//	//// 用select可以更新零值
//	//DB.Model(&Student2{}).Where("age = ?", 21).Select("gender", "address").Updates(Student2{
//	//	Address: &address,
//	//	Gender:  false,
//	//})
//	//
//	//// 简洁版 使用map
//	//DB.Model(&Student2{}).Where("age = ?", 21).Updates(map[string]any{
//	//	"address": address,
//	//	"gender":  false,
//	//})
//
//	//// 删除 delete
//	//DB.Delete(&student, []int{12, 13})
//	//
//	//DB.Take(&student)
//	//DB.Delete(&student)
//
//	DB.Create(&Student2{
//		Name: "lwk666",
//		Age:  0,
//	})
//}
