/*
 * @Description: Do not edit
 * @Date: 2022-02-13 15:31:33
 * @LastEditors: wanghaijie01
 * @LastEditTime: 2022-02-13 15:37:03
 */

package main

import (
	"context"
	"fmt"
	"time"
)

// 利用ctx传递取消信号
func main() {
	parentCtx := context.Background()
	ctx, cancel := context.WithCancel(parentCtx)

	go func() {
		time.Sleep(1 * time.Second)
		cancel()
	}()

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("in goroutine 1, ctx is canceled!", ctx.Err())
				return
			}
		}
	}()

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("in goroutine 2, ctx is canceled!")
				return
			}
		}
	}()

	time.Sleep(3 * time.Second)
}
