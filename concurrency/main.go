/*
 * @Description: Do not edit
 * @Date: 2022-01-14 00:55:54
 * @LastEditors: wanghaijie01
 * @LastEditTime: 2022-01-14 01:33:36
 */

package main

import (
	"fmt"
	"runtime"
	"time"
)

// 控制并发数
// 利用channel控制任务执行的并发度为n
func concurrencyN(n int) {
	limit := make(chan struct{}, n)
	// for 循环表示有无穷无尽的工作需要处理
	for {
		go func() {
			limit <- struct{}{}
			// 用sleep模拟要处理的工作
			time.Sleep(2 * time.Second)
			<-limit
		}()
	}
}

func main() {
	go concurrencyN(100)
	for {
		fmt.Println("num of goroutine: ", runtime.NumGoroutine())
		time.Sleep(time.Second)
	}
}
