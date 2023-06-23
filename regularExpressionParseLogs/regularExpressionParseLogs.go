package main

import (
	"fmt"
	"regexp"
)

var (
	reg  = regexp.MustCompile(`use time (\d+)ms`)
	regC *regexp.Regexp
)

// init()函数通常用于全局变量的初始化
func init() {
	var err error
	if regC, err = regexp.Compile(`use time (\d+)ms`); err != nil {
		panic(err)
	}
}

func main() {
	log := "recall use time 38ms, sort use time 20ms"    //这个字符串里reg模式命中两次
	indexs1 := reg.FindAllSubmatchIndex([]byte(log), -1) //-1表示返回所有匹配上reg的地方
	fmt.Println(indexs1)
	indexs2 := reg.FindAllSubmatchIndex([]byte(log), 1) //1表示只需要返回1处(最靠前的1处)匹配上reg的地方
	fmt.Println(indexs2)
	subMatch := indexs1[0]
	begin, end := subMatch[0], subMatch[1]
	//整体匹配上reg的部分
	fmt.Println(log[begin:end])
	begin, end = subMatch[2], subMatch[3]
	//匹配上reg中()的部分
	fmt.Println(log[begin:end])
}

//regexp.MustCompile发生异常时会自动触发panic，这里使用regexp.Compile，发生异常时自行调用了panic。
