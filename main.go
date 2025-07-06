package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("生产者进程启动")

	file, err := os.OpenFile("data.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	for i := 1; i <= 5; i++ {
		data := fmt.Sprintf("数据-%d\n", i)
		file.WriteString(data)
		fmt.Printf("生产者写入: %s", data)
		time.Sleep(2 * time.Second)
	}
}
