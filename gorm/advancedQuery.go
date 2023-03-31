package main

import (
	"fmt"
	"gorm.io/gorm"
)

type Student2 struct {
	ID     uint    `gorm:"size:10"` // 默认作为主键
	Name   string  `gorm:"type:varchar(16)"`
	Age    int     `gorm:"size:3"`
	Email  *string `gorm:"size:128"` //使用指针可以存储空值
	Gender bool    `gorm:"default:true"`
}

func PtrString(email string) *string {
	return &email
}

// model层写一些通用的查询方式，这样外界就可以直接调用方法即可
func Age23(db *gorm.DB) *gorm.DB {
	return db.Where("age > ?", 23)
}

func main() {
	//var studentList []Student2
	//DB.Debug().AutoMigrate(&Student2{})
	//DB.Find(&studentList).Delete(&studentList)
	//studentList = []Student2{
	//	{ID: 1, Name: "李元芳", Age: 32, Email: PtrString("lyf@yf.com"), Gender: true},
	//	{ID: 2, Name: "张武", Age: 18, Email: PtrString("zhangwu@lly.cn"), Gender: true},
	//	{ID: 3, Name: "枫枫", Age: 23, Email: PtrString("ff@yahoo.com"), Gender: true},
	//	{ID: 4, Name: "刘大", Age: 54, Email: PtrString("liuda@qq.com"), Gender: true},
	//	{ID: 5, Name: "李武", Age: 23, Email: PtrString("liwu@lly.cn"), Gender: true},
	//	{ID: 6, Name: "李琦", Age: 14, Email: PtrString("liqi@lly.cn"), Gender: false},
	//	{ID: 7, Name: "晓梅", Age: 25, Email: PtrString("xiaomeo@sl.com"), Gender: false},
	//	{ID: 8, Name: "如燕", Age: 26, Email: PtrString("ruyan@yf.com"), Gender: false},
	//	{ID: 9, Name: "魔灵", Age: 21, Email: PtrString("moling@sl.com"), Gender: true},
	//}
	//DB.Create(&studentList)

	//// Where
	//// 等价于sql语句中的where
	//var users []Student2
	//// 查询用户名是枫枫的
	//DB.Where("name = ?", "枫枫").Find(&users)
	//fmt.Println(users)
	//// 查询用户名不是枫枫的
	//DB.Where("name <> ?", "枫枫").Find(&users)
	//fmt.Println(users)
	//// 查询用户名包含 如燕，李元芳的
	//DB.Where("name in ?", []string{"如燕", "李元芳"}).Find(&users)
	//fmt.Println(users)
	//// 查询姓李的
	//DB.Where("name like ?", "李%").Find(&users)
	//fmt.Println(users)
	//// 查询年龄大于23，是qq邮箱的
	//DB.Where("age > ? and email like ?", "23", "@qq.com").Find(&users)
	//fmt.Println(users)
	//// 查询是qq邮箱的，或者是女的
	//DB.Where("gender = ? or email like ?", "false", "@qq.com").Find(&users)
	//fmt.Println(users)
	//
	//// 使用结构体查询，会过滤零值，结构体中的条件之间为and关系
	//DB.Where(&Student2{Name: "李元芳", Age: 0}).Find(&users)
	//fmt.Println(users)
	//
	//// 使用map查询，不会过滤零值
	//DB.Where(map[string]any{"name": "李元芳", "age": 0}).Find(&users)
	//fmt.Println(users)

	//// Not条件，与where中的not等价
	//// 排除年龄大于23的
	//DB.Not("age > 23").Find(&users)
	//fmt.Println(users)
	//
	//// Or条件，与where中的or等价
	//DB.Or("gender = ?", false).Or("email like ?", "%@qq.com").Find(&users)
	//fmt.Println(users)

	//// Select选择字段
	//DB.Select("name", "age").Find(&users)
	//fmt.Println(users)
	//
	//// 可以使用Scan。将选择的字段存入另一个结构体中
	//type User struct {
	//	Name string
	//	Age  int
	//}
	//var students []Student2
	//var users []User
	//// 从students中筛选出两个字段，存储到users结构体实例化的对象中
	//// 但是这种写法会查询两次
	//// 因为是先找到，再访问存储到另一个地方
	//DB.Select("name", "age").Find(&students).Scan(&users)
	//fmt.Println(users)
	//// 这种写法就只查询一次
	//// 因为model选出字段直接存储
	//DB.Model(&Student2{}).Select("name", "age").Scan(&users)
	//fmt.Println(users)
	//// 这种也是查询 一次
	//// 相对于上面那种Model，这种是对于Model实例化的表进行筛选，剩下同理
	//DB.Table("f_student2").Select("name", "age").Scan(&users)
	//fmt.Println(users)
	//// Scan是根据column列名进行扫描的
	//type User2 struct {
	//	Name123 string `gorm:"column:name"`
	//	Age     int
	//}
	//var users2 []User2
	//DB.Table("f_student2").Select("name", "age").Scan(&users)
	//fmt.Println(users2)

	//var users []Student2
	//// 排序
	//DB.Order("age desc").Find(&users)
	//fmt.Println(users)

	//// 分页查询
	//// 一页两条，第一页
	//DB.Limit(2).Offset(0).Find(&users)
	//fmt.Println(users)
	//// 第二页
	//DB.Limit(2).Offset(2).Find(&users)
	//fmt.Println(users)
	//// 第三页
	//DB.Limit(2).Offset(4).Find(&users)
	//fmt.Println(users)
	//// 通用写法
	//// 一页多少条
	//limit := 2
	//// 第几页
	//page := 1
	//offset := (page - 1) * limit
	//DB.Limit(limit).Offset(offset).Find(&users)
	//fmt.Println(users)

	//// 去重
	//var ageList []int
	//DB.Table("f_student2").Select("age").Distinct("age").Scan(&ageList)
	////// 或者
	////DB.Table("f_student2").Select("distinct age").Scan(&ageList)
	//fmt.Println(ageList)

	//// 分组查询
	////var ageList []int
	////// 查询男生的个数和女生的个数
	////DB.Table("f_student2").Select("count(id)").Group("gender").Scan(&ageList)
	////fmt.Println(ageList)
	//
	////// 精确到哪一个是男生个数，哪一个是女生个数
	////type AggeGroup struct {
	////	Gender int
	////	Count  int `gorm:"column:count(id)"`
	////}
	////var agge []AggeGroup
	////// 查询男生的个数和女生的个数
	////DB.Table("f_student2").Select("count(id)", "gender").Group("gender").Scan(&agge)
	////fmt.Println(agge)
	//
	//// 再精确到具体的男女名字
	//type AggeGroup struct {
	//	Gender int
	//	Count  int    `gorm:"column:count(id)"`
	//	Name   string `gorm:"column:group_concat(name)"`
	//}
	//var agge []AggeGroup
	//// 查询男生的个数和女生的个数
	//DB.Table("f_student2").Select("count(id)", "gender", "group_concat(name)").Group("gender").Scan(&agge)
	//fmt.Println(agge)
	//
	//// 直接执行原生SQL
	//DB.Raw(`SELECT count(id), gender, group_concat(name) FROM f_student2 GROUP BY gender`).Scan(&agge)
	//fmt.Println(agge)

	// 子查询
	var users []Student2
	//// 查询大于平均年龄的用户
	//// 原生sql
	//DB.Raw(`select * from f_student2 where age > (select avg(age) from f_student2)`).Scan(&users)
	//// gorm
	//DB.Model(Student2{}).Where("age > (?)", DB.Model(Student2{}).Select("avg(age)")).Find(&users)
	//fmt.Println(users)

	//// 命名参数
	//DB.Where("name = @name and age = @age", sql.Named("name", "枫枫"), sql.Named("age", 23)).Find(&users)
	//DB.Where("name = @name and age = @age", map[string]any{"name": "枫枫", "age": 23}).Find(&users)
	//fmt.Println(users)

	//// find到map
	//var res []map[string]any
	//DB.Table("f_student2").Find(&res)
	//fmt.Println(res)

	// model层写一些通用的查询方式，这样外界就可以直接调用方法即可
	// 查询引用Scope
	DB.Scopes(Age23).Find(&users)
	fmt.Println(users)
}
