package main

import (
	"fmt"
	"sync"
)

func Word(id int) {
	fmt.Printf("第%d个goroutine在工作\n", id)
	wg.Done()
}

var wg sync.WaitGroup

func main() {
	//go fmt.Println("goroutine1")
	//go test()
	//GMP调度模型
	//go func() {
	//	fmt.Println("goroutine3")
	//}() //匿名函数末尾的括号表示自调用
	//fmt.Println("limeng")
	//time.Sleep(1 * time.Second)
	wg.Add(5)
	for i := 1; i <= 5; i++ {
		go Word(i)
	}
	//time.Sleep(1 * time.Second)
	wg.Wait()
}
