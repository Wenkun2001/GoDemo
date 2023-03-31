package main

// User 一个用户可以发布多篇文章，一篇文章属于一个用户
type User struct {
	ID       uint
	Name     string    `gorm:"size:8"`
	Articles []Article // 用户拥有的文章列表
}

// Article 外键名称是关联表名+ID，类型是uint
type Article struct {
	ID     uint
	Title  string `gorm:"size:16"`
	UserID uint   // 属于   这里的类型要和引用的外键类型一致，包括大小
	User   User   // 属于
}

//// 重写外键关联
//// 将UID作为Article的外键，User外键关系要指向UID
//// User所拥有的Articles也得更改外键，改为UID
//type User struct {
//	ID       uint
//	Name     string    `gorm:"size:8"`
//	Articles []Article `gorm:"foreignKey:UID"` // 用户拥有的文章列表
//}
//
//type Article struct {
//	ID    uint
//	Title string `gorm:"size:16"`
//	UID   uint   // 属于
//	User  User   `gorm:"foreignKey:UID"` // 属于
//}
//
//// 重写外键引用
//// 直接关联Name，原本与user_id的关系变为与user_name的关系
//type User struct {
//	ID       uint
//	Name     string    `gorm:"size:8"`
//	Articles []Article `gorm:"foreignKey:UserName;references:Name"` // 用户拥有的文章列表
//}
//
//type Article struct {
//	ID       uint
//	Title    string `gorm:"size:16"`
//	UserName string `gorm:"size:8"`
//	User     User `gorm:"references:Name"` // 属于
//}

func main() {
	//DB.Debug().AutoMigrate(&User{}, &Article{})
	// 一对多的添加
	//// 创建用户，并且创建文章
	//// gorm自动创建了两篇文章，以及创建了一个用户，还将他们的关系给关联上了
	//a1 := Article{Title: "python"}
	//a2 := Article{Title: "goland"}
	//user := User{Name: "枫枫", Articles: []Article{a1, a2}}
	//DB.Create(&user)

	//// 创建文章，关联已有账户
	//// 方法一
	//a1 := Article{Title: "golang零基础入门", UserID: 1}
	//DB.Create(&a1)
	//// 方法二
	//var user User
	//DB.Take(&user, 1)
	//DB.Create(&Article{Title: "python进阶", User: user})
	//// 方法三
	//DB.Create(&Article{Title: "pytorch", User: User{Name: "lwk"}})

	//// 外键添加
	//var user User
	//DB.Take(&user, 2)
	//var article Article
	//DB.Take(&article, 5)
	//
	//// 给现有用户绑定文章
	//user.Articles = []Article{article}
	//DB.Save(&user)
	//// Append方法
	//DB.Model(&user).Association("Articles").Append(&article)
	//
	//// 给现有文章关联用户
	//article.UserID = 2
	//DB.Save(&article)
	//// Append方法
	//DB.Model(&article).Association("User").Append(&user)

	// 查询
	// 查询用户，显示用户的文章列表
	// 这样是无法显示出完整文章列表的
	var user User
	//DB.Take(&user, 1)
	//fmt.Println(user)

	//// 预加载
	//// 预加载的名字就是外键关联的属性名
	//// 使用预加载，来加载需要查询用户的文章列表
	//DB.Preload("Articles").Take(&user, 1)
	//fmt.Println(user)
	//
	//// 查询文章，使用预加载，显示文章用户的信息
	//var article Article
	//DB.Preload("User").Take(&article, 1)
	//fmt.Println(article)
	//
	// 嵌套预加载
	// 查询文章，显示用户，并且显示用户关联的所有文章
	// 即在上面显示用户的基础上，还显示用户关联的所有文章
	//DB.Preload("User.Articles").Take(&article, 1)
	//fmt.Println(article)
	//
	// 带条件的预加载
	// 查询用户的所有文章列表，过滤文章，只有id为1的文章被预加载出来
	//DB.Preload("Articles", "id = ?", 1).Take(&user, 1)
	//fmt.Println(user)
	//
	//// 自定义预加载
	//DB.Preload("Article", func(db *gorm.DB) *gorm.DB {
	//	return db.Where("id in ?", []int{1, 2})
	//}).Take(&user, 1)
	//fmt.Println(user)

	//// 删除
	//// 级联删除
	//// 删除用户，与用户关联的文章也会删除
	//DB.Take(&user, 1)
	//DB.Select("Articles").Delete(&user)

	// 清除外键关系
	// 删除用户，与将与用户关联的文章，外键设置为null
	DB.Preload("Articels").Take(&user, 2)
	DB.Model(&user).Association("Articles").Delete(&user.Articles)

}
