package main

import (
	"fmt"
	"time"
)

// ArticleModel 自定义连接表主键
// 这个功能还是很有用的，例如文章表可能叫ArticleModel，标签表可能叫TagModel
// 那么按照gorm默认的主键名，那就分别是ArticleModelID，TagModelID，太长了，根本就不实用
// 主要是要修改这两项
// joinForeignKey 连接的主键id
// JoinReferences 关联的主键id
type ArticleModel struct {
	ID    uint
	Title string
	Tags  []TagModel `gorm:"many2many:article_tags;joinForeignKey:ArticleID;JoinReferences:TagID"`
}

type TagModel struct {
	ID       uint
	Name     string
	Articles []ArticleModel `gorm:"many2many:article_tags;joinForeignKey:TagID;JoinReferences:ArticleID"`
}

type ArticleTagModel struct {
	ArticleID uint `gorm:"primaryKey"` // article_id
	TagID     uint `gorm:"primaryKey"` // tag_id
	CreatedAt time.Time
}

// UserModel 操作连接表
// 如果通过一张表去操作连接表，这样会比较麻烦
// 比如查询某篇文章关联了哪些标签
// 或者是举个更通用的例子，用户和文章，某个用户在什么时候收藏了哪篇文章
// 无论是通过用户关联文章，还是文章关联用户都不太好查
// 最简单的就是直接查连接表
type UserModel struct {
	ID       uint
	Name     string
	Collects []ArticleModel2 `gorm:"many2many:user_collect_models;joinForeignKey:UserID;JoinReferences:ArticleID"`
}

type ArticleModel2 struct {
	ID    uint
	Title string
	// 这里也可以反向引用，根据文章查哪些用户收藏了
}

// UserCollectModel 用户收藏文章表
type UserCollectModel struct {
	UserID    uint `gorm:"primaryKey"` // article_id
	ArticleID uint `gorm:"primaryKey"` // tag_id
	CreatedAt time.Time
}

// UserModel2 修改表结构，不需要重新迁移，加一些字段
type UserModel2 struct {
	ID       uint
	Name     string
	Collects []ArticleModel3 `gorm:"many2many:user_collect_models;joinForeignKey:UserID;JoinReferences:ArticleID"`
}

type ArticleModel3 struct {
	ID    uint
	Title string
}

// UserCollectModel2 用户收藏文章表
type UserCollectModel2 struct {
	UserID        uint          `gorm:"primaryKey"` // article_id
	UserModel2    UserModel2    `gorm:"foreignKey:UserID"`
	ArticleID     uint          `gorm:"primaryKey"` // tag_id
	ArticleModel3 ArticleModel3 `gorm:"foreignKey:ArticleID"`
	CreatedAt     time.Time
}

func main() {
	// SetupJoinTable
	// 添加和更新时用这个
	// 这样才能走自定义的连接表，以及走它的钩子函数
	// 查询则不需要这个
	//
	// 自定义连接表主键
	//DB.SetupJoinTable(&ArticleModel{}, "Tags", &ArticleTagModel{})
	//DB.SetupJoinTable(&TagModel{}, "Articles", &ArticleTagModel{})
	//DB.Debug().AutoMigrate(&ArticleModel{}, &TagModel{}, &ArticleTagModel{})
	// 操作同自定义连接表

	// 操作连接表
	//DB.SetupJoinTable(&UserModel{}, "Collects", &UserCollectModel{})
	//DB.Debug().AutoMigrate(&UserModel{}, &ArticleModel2{}, &UserCollectModel{})
	//
	//// 常用的操作就是根据用户查收藏的文章列表
	//// 不好分页，并且也拿不到收藏文章的时间
	//var user UserModel
	//DB.Preload("Collects").Take(&user, "name = ?", "枫枫")
	//fmt.Println(user)
	//
	//// 这样虽然可以查到用户id，文章id，收藏的时间
	//// 但是搜索只能根据用户id，返回也拿不到用户名，文章标题等
	//var collects []UserCollectModel
	//DB.Find(&collects, "user_id = ?", 2)
	//fmt.Println(collects)

	// 修改表结构，不需要重新迁移，加一些字段
	//DB.SetupJoinTable(&UserModel2{}, "Collects", &UserCollectModel2{})
	//DB.Debug().AutoMigrate(&UserModel2{}, &ArticleModel3{}, &UserCollectModel2{})
	// 查询
	var collects []UserCollectModel
	var user UserModel
	DB.Take(&user, "name = ?", "枫枫")
	// 这里用map的原因是如果没查到，那就会查0值，如果是struct，则会忽略零值，全部查询
	DB.Debug().Preload("UserModel").Preload("ArticleModel").Where(map[string]any{"user_id": user.ID}).Find(&collects)
	for _, collect := range collects {
		fmt.Println(collect)
	}
}
