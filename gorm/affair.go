package main

import (
	"fmt"
	"gorm.io/gorm"
)

type User5 struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Money int    `json:"money"`
}

// InnoDB引擎才支持事务，MyISAM不支持事务
// DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{})

func main() {
	//DB.Debug().AutoMigrate(&User5{})
	var zhangsan, lisi User5
	DB.Take(&zhangsan, "name = ?", "张三")
	DB.Take(&lisi, "name = ?", "李四")

	// 普通事务
	// 张三给李四转账100元
	DB.Transaction(func(tx *gorm.DB) error {
		// 先给张三-100
		zhangsan.Money -= 100
		err := tx.Model(&zhangsan).Update("money", zhangsan.Money).Error
		if err != nil {
			fmt.Println(err)
			return err
		}
		// 再给李四+100
		lisi.Money += 100
		err = tx.Model(&lisi).Update("money", lisi.Money).Error
		if err != nil {
			fmt.Println(err)
			return err
		}
		// 提交事务
		return nil
	})

	// 手动事务
	// 张三给李四转账100元
	tx := DB.Begin()
	// 先给张三-100
	zhangsan.Money -= 100
	err := tx.Model(&zhangsan).Update("money", zhangsan.Money).Error
	if err != nil {
		tx.Rollback()
	}
	// 再给李四+100
	lisi.Money += 100
	err = tx.Model(&lisi).Update("money", lisi.Money).Error
	if err != nil {
		tx.Rollback()
	}
	// 提交事务
	tx.Commit()
}
