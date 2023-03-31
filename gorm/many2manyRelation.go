package main

import (
	"time"
)

// Tag 多对多关系
// 需要用第三张表存储两张表的关系
type Tag struct {
	ID       uint
	Name     string
	Articles []Article2 `gorm:"many2many:article_tags;"` // 用于反向引用
}

type Article2 struct {
	ID    uint
	Title string
	Tags  []Tag `gorm:"many2many:article_tags;"`
}

type Article3 struct {
	ID    uint
	Title string
	Tags  []Tag2 `gorm:"many2many:article_tags;"`
}

type Tag2 struct {
	ID   uint
	Name string
}

type ArticleTag struct {
	Article3ID uint `gorm:"primaryKey"`
	Tag2ID     uint `gorm:"primaryKey"`
	CreatedAt  time.Time
}

func main() {
	//DB.Debug().AutoMigrate(&Tag{}, &Article2{})
	//DB.Create(&Article2{
	//	Title: "python基础课程",
	//	Tags: []Tag{
	//		{Name: "python"},
	//		{Name: "基础课程"},
	//	},
	//})

	//// 添加文章，选择标签
	//var tags []Tag
	//DB.Find(&tags, "name = ?", "基础课程")
	//DB.Create(&Article2{
	//	Title: "golang基础",
	//	Tags:  tags,
	//})

	// 多对多查询
	// 查询文章，显示文章的标签列表
	//var article Article2
	//DB.Preload("Tags").Take(&article, 1)
	//fmt.Println(article)
	// 查询标签，显示文章列表
	//var tag Tag
	//DB.Preload("Articles").Take(&tag, 2)
	//fmt.Println(tag)

	//// 多对多更新
	//// 移除文章的标签
	//DB.Preload("Tags").Take(&article, 1)
	//DB.Model(&article).Association("Tags").Delete(&article.Tags)
	//fmt.Println(article)
	//// 更新文章的标签
	//var tags []Tag
	//DB.Find(&tag, []int{2, 6, 7})
	//DB.Preload("Tags").Take(&article, 1)
	//DB.Model(&article).Association("Tags").Replace(tags)
	//fmt.Println(article)

	//// 设置Article的Tags表为ArticleTag
	//DB.SetupJoinTable(&Article3{}, "Tags", &ArticleTag{})
	//// 如果tag要反向应用Article，那么也得加上
	//DB.Debug().AutoMigrate(&Article3{}, &Tag2{}, &ArticleTag{})

	DB.SetupJoinTable(&Article3{}, "Tags", &ArticleTag{}) // 要设置这个，才能走到我们自定义的连接表
	//DB.Debug().AutoMigrate(&Article3{}, &Tag2{}, &ArticleTag{})
	//// 1、添加文章并添加标签，并自动关联
	//DB.Create(&Article3{
	//	Title: "flask零基础入门",
	//	Tags: []Tag2{
	//		{Name: "python"},
	//		{Name: "后端"},
	//		{Name: "web"},
	//	},
	//	// CreatedAt time.Time 由于我们设置的是CreatedAt，gorm会自动填充当前时间，
	//	// 如果是其他的字段，需要使用到ArticleTag 的添加钩子 BeforeCreate
	//})
	//
	//// 2、添加文章，关联已有标签
	var tags []Tag2
	//DB.Find(&tags, "name in ?", []string{"python", "web"})
	//DB.Create(&Article3{
	//	Title: "flask请求对象",
	//	Tags:  tags,
	//})
	//
	//// 3、给已有文章关联标签
	//article := Article3{
	//	Title: "django基础",
	//}
	//DB.Create(&article)
	//var at Article3
	//DB.Find(&tags, "name in ?", []string{"python", "web"})
	//DB.Take(&at, article.ID).Association("Tags").Append(tags)
	//
	// 4、替换已有文章的标签
	var article Article3
	DB.Find(&tags, "name in ?", []string{"后端"})
	DB.Take(&article, "title = ?", "flask请求对象")
	DB.Model(&article).Association("Tags").Replace(tags)
	//
	//// 5、查询文章列表，显示标签
	//var articles []Article3
	//DB.Preload("Tags").Find(&articles)
	//fmt.Println(articles)
}
