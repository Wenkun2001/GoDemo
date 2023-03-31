package main

import (
	"encoding/json"
	"fmt"
	"time"
)

// 结构体怎么表示数据库中的字段呢？
// golang中没有枚举
// 我们只能自己通过逻辑实现枚举

type Weekday int

const (
	Sunday    Weekday = iota + 1 // EnumIndex = 1
	Monday                       // EnumIndex = 2
	Tuesday                      // EnumIndex = 3
	Wednesday                    // EnumIndex = 4
	Thursday                     // EnumIndex = 5
	Friday                       // EnumIndex = 6
	Saturday                     // EnumIndex = 7
)

var WeekStringList = []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
var WeekTypeList = []Weekday{Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday}

// String 转字符串
func (w Weekday) String() string {
	return WeekStringList[w-1]
}

// MarshalJSON 自定义类型转换为json
func (w Weekday) MarshalJSON() ([]byte, error) {
	return json.Marshal(w.String())
}

// EnumIndex 自定义类型转原始类型
func (w Weekday) EnumIndex() int {
	return int(w)
}

// ParseWeekDay 字符串转自定义类型
func ParseWeekDay(week string) Weekday {
	// 字符串与WeekStringList字符串数组里面数据遍历对比
	// 相同则通过index从WeekTypeList数组中获取相对应的自定义类型数据
	for i, i2 := range WeekStringList {
		if week == i2 {
			return WeekTypeList[i]
		}
	}
	return Monday
}

// ParseIntWeekDay 数字转自定义类型
func ParseIntWeekDay(week int) Weekday {
	return Weekday(week)
}

type DayInfo struct {
	Weekday Weekday   `json:"weekday"`
	Date    time.Time `json:"date"`
}

func main() {
	w := Sunday
	fmt.Println(w)
	dayInfo := DayInfo{Weekday: Sunday, Date: time.Now()}
	// 自定义类型转换为json
	data, err := json.Marshal(dayInfo)
	fmt.Println(string(data), err)
	// 字符串转自定义类型
	week := ParseWeekDay("Sunday")
	fmt.Println(week)
	// 数字转自定义类型
	week = ParseIntWeekDay(2)
	fmt.Println(week)
}

// 在需要输出的时候（print，json），自定义类型就变成了字符串
// 从外界接收的数据也能转换为自定义类型，这就是golang中的枚举，假枚举
