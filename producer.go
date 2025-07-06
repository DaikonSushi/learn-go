package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("生产者程序启动")
	for i := 1; i <= 10; i++ {
		fmt.Printf("生产者生成数据: %d\n", i)
		time.Sleep(1 * time.Second)
	}
	fmt.Println("生产者程序结束")
}
