/*
 * @Description: Do not edit
 * @Date: 2022-01-21 15:25:56
 * @LastEditors: wanghaijie01
 * @LastEditTime: 2022-01-22 00:07:59
 */
// nil 结构体转 interface，不等于 nil
// interface 等于 nil 条件是
package main

import (
	"fmt"
)

// Student xxx
type Student struct {
	Name string
}

// Rename xxx
func (s *Student) Rename(name string) {
	s.Name = name
}

// People xxx
type People interface {
	Rename(string)
}

func main() {
	var s *Student
	fmt.Println("s == nil:", s == nil)
	p := People(s)
	fmt.Println("p == nil:", p == nil)
}
