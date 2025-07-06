package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("消费者程序启动")
	for i := 1; i <= 10; i++ {
		fmt.Printf("消费者处理数据: %d\n", i)
		time.Sleep(1500 * time.Millisecond)
	}
	fmt.Println("消费者程序结束")
}
