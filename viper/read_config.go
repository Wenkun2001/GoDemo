package main

import (
	"fmt"
	"github.com/spf13/viper"
)

// 使用相同的代码可以解析JSON, TOML, YAML, HCL, envfile and Java properties等多种格式的配置文件。
// 以JSON和YAML举例演示
func readContent(config *viper.Viper) {
	if err := config.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("找不到配置文件..")
		} else {
			fmt.Printf("解析配置文件出错: %v\n", err)
		}
	}

	// 读取配置
	user1 := config.GetString("section1.user")
	user2 := config.GetString("section2.user")
	height := config.GetInt32("section1.body.height")
	weight := config.GetInt32("section1.body.weight")
	fmt.Println(user1, user2, height, weight)
}

func readJson() {
	config := viper.New()
	// 文件所在目录
	config.AddConfigPath("file/conf/")
	// 文件名
	config.SetConfigName("account")
	// 文件类型
	config.SetConfigType("json")

	readContent(config)
}

func readYaml() {
	config := viper.New()
	// 文件所在目录
	config.AddConfigPath("file/conf/")
	// 文件名
	config.SetConfigName("account")
	// 文件类型
	config.SetConfigType("yaml")

	readContent(config)
}

func main() {
	readJson()
	readYaml()
}
