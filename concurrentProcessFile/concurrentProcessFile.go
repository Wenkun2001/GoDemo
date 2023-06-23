package main

import (
	"bufio"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

var (
	regex *regexp.Regexp
	sum   int32
)

func init() {
	var err error
	if regex, err = regexp.Compile(`recall (\d+) by radic, recall \d+ by milvus_short, recall \d+ by milvus_long`); err != nil {
		panic(err)
	}
}

// ExtractNumber 成功匹配上则返回非负整数，否则返回-1
func ExtractNumber(log string) int {
	indexs := regex.FindAllSubmatchIndex([]byte(log), -1)
	if len(indexs) > 0 {
		subMatch := indexs[0]
		begin, end := subMatch[2], subMatch[3]
		match := log[begin:end]
		if n, err := strconv.Atoi(match); err != nil {
			return -1
		} else {
			return n
		}
	} else {
		return -1
	}
}

// ListDir 递归遍历目录，返回下面的所有文件（不包含文件夹）
func ListDir(dir string) []string {
	files := make([]string, 0, 20)
	//WalkDir内部会递归遍历目录
	filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error { //WalkDir比Walk更高效
		if err != nil {
			return err
		} else if info, err := d.Info(); err == nil {
			if info.Mode().IsRegular() {
				files = append(files, path)
			}
			return nil
		} else {
			return err
		}
	})
	fmt.Printf("%s目录下共%d个文件\n", dir, len(files))
	return files
}

// ProcessFile 处理文件
func ProcessFile(file string) {
	fin, err := os.Open(file)
	if err != nil {
		fmt.Printf("打开文件失败:%v", err)
		return
	}
	defer fin.Close()

	reader := bufio.NewReader(fin)
	for {
		if log, err := reader.ReadString('\n'); err != nil {
			if err == io.EOF {
				if len(log) > 0 {
					n := ExtractNumber(log)
					if n >= 0 {
						// sum += int32(n)
						atomic.AddInt32(&sum, int32(n)) //并发安全
					}
				}
			} else {
				fmt.Printf("读文件失败:%v", err)
			}
			break
		} else {
			log = strings.TrimRight(log, "\n")
			if len(log) > 0 {
				n := ExtractNumber(log)
				if n >= 0 {
					// sum += int32(n)
					atomic.AddInt32(&sum, int32(n)) //并发安全
				}
			}
		}
	}
}

// ProcessDir 处理目录
func ProcessDir(dir string) {
	files := ListDir(dir)
	for _, file := range files {
		ProcessFile(file)
	}
}

// ConcurrentProcessDir 并发处理目录
func ConcurrentProcessDir(dir string) {
	files := ListDir(dir)
	wg := sync.WaitGroup{}
	wg.Add(len(files))
	for _, file := range files {
		go func(file string) {
			defer wg.Done()
			ProcessFile(file)
		}(file) //协程中使用for循环生成的变量时，务必把变量拷贝到协程里去
	}
	wg.Wait()
}

func main() {
	dir := "file/data/log" //日志文件存放的路径，file是相对于执行go run的路径
	begin := time.Now()
	// ProcessDir(dir)  //串行耗时8.8秒。耗时跟电脑有关
	ConcurrentProcessDir(dir) //并行耗时3.8秒。耗时跟电脑有关
	fmt.Printf("sum=%d, time %dms\n", sum, time.Since(begin).Milliseconds())
}
