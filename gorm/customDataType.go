package main

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

// 自定义的数据类型必须实现 Scanner 和 Valuer 接口
// 让gorm知道如何将该类型接收、保存到数据库

// Info 存储结构体
type Info struct {
	Status string `json:"status"`
	Addr   string `json:"addr"`
	Age    int    `json:"age"`
}

// Scan 从数据库中读取出来
func (i *Info) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}
	info := Info{}
	err := json.Unmarshal(bytes, &info)
	*i = info
	return err
}

// Value 存入数据库
func (i Info) Value() (driver.Value, error) {
	return json.Marshal(i)
}

type User4 struct {
	ID   uint
	Name string
	Info Info `gorm:"type:string"`
}

// 枚举类型

// Host 枚举1.0
// 很多时候，我们会对一些状态进行判断，而这些状态都是有限的
// 例如，主机管理中，状态有 Running 运行中， OffLine 离线， Except 异常
// 如果存储字符串，不仅是浪费空间，每次判断还要多复制很多字符，最主要是后期维护麻烦
type Host struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Status int    `json:"status"`
}

// 用常量存储不变的值
// 虽然代码变多了，但是维护方便了
// 但是数据库中存储的依然是字符串，浪费空间这个问题并没有解决
//const (
//	Running = "Running"
//	Except  = "Except"
//	OffLine = "OffLine"
//)

// 枚举2.0
// 使用数字表示状态
// 返回数据给前端，前端接收到的状态就是数字
const (
	Running = 1
	Except  = 2
	OffLine = 3
)

// MarshalJSON 在json序列化的时候，根据映射转换回去
// 这样写确实可以实现我们的需求，但是根本就不够通用，凡是用到枚举，都得给这个Struct实现MarshalJSON方法
func (h Host) MarshalJSON() ([]byte, error) {
	var status string
	switch h.Status {
	case Running:
		status = "Running"
	case Except:
		status = "Except"
	case OffLine:
		status = "OffLine"
	}
	return json.Marshal(&struct {
		ID     uint   `json:"id"`
		Name   string `json:"name"`
		Status string `json:"status"`
	}{
		ID:     h.ID,
		Name:   h.Name,
		Status: status,
	})
}

type Status int

// MarshalJSON 枚举3.0
func (status Status) MarshalJSON() ([]byte, error) {
	var str string
	switch status {
	case Running:
		str = "Running"
	case Except:
		str = "Except"
	case OffLine:
		str = "Status"
	}
	return json.Marshal(str)
}

func main() {
	//DB.Debug().AutoMigrate(&Info{}, &User4{})
	//// 添加
	//DB.Create(&User4{
	//	Name: "枫枫",
	//	Info: Info{
	//		Status: "牛逼",
	//		Addr:   "成都市",
	//		Age:    21,
	//	},
	//})
	//// 查询
	//var user User4
	//DB.Take(&user)
	//fmt.Println(user)

	//// 枚举1.0
	//host := Host{}
	//if host.Status == Running {
	//	fmt.Println("在线")
	//}
	//if host.Status == Except {
	//	fmt.Println("异常")
	//}
	//if host.Status == OffLine {
	//	fmt.Println("离线")
	//}

	// 枚举2.0 3.0
	// 在json序列化的时候，根据映射转换回去
	host := Host{1, "枫枫", Running}
	data, _ := json.Marshal(host)
	fmt.Println(string(data)) // {"id":1,"name":"枫枫","status":"Running"}
}
