/*
 * @Description: Do not edit
 * @Date: 2022-02-14 00:18:04
 * @LastEditors: wanghaijie01
 * @LastEditTime: 2022-02-14 15:34:13
 */

package main

import (
	"fmt"
	"reflect"
)

// P xxx
type P struct {
	Name string
}

// Q xxx
type Q struct {
	Name string
}

func main() {
	p := P{}
	pointP := reflect.ValueOf(&p)
	em := pointP.Elem()
	fmt.Println("em can set:", em.CanSet())
	name := em.FieldByName("Name")
	name.SetString("wanghaijie")
	fmt.Println("new p is ", p)
}
