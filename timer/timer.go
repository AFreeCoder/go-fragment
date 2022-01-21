/*
 * @Description: Do not edit
 * @Date: 2022-01-20 00:42:24
 * @LastEditors: wanghaijie01
 * @LastEditTime: 2022-01-20 13:20:10
 */

package main

import (
	"fmt"
	"time"
)

// doWorkwithTicker  使用ticker作为定时器
func doWorkwithTicker() {
	ticker := time.Tick(10 * time.Second)
	for {
		select {
		case <-ticker:
			fmt.Println("执行定时任务...")
		}
	}
}

// doWorkwithTimer  使用timer作为定时器
func doWorkwithTimer() {
	timer := time.NewTimer(10 * time.Second)
	for {
		select {
		case n := <-timer.C:
			fmt.Println("执行定时任务...", n)
			timer.Reset(10 * time.Second)  // 这里注意需要reset
		}
	}
}

func main() {
	doWorkwithTicker()
	doWorkwithTimer()
}
