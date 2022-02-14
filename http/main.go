/*
 * @Description: Do not edit
 * @Date: 2022-01-21 19:42:09
 * @LastEditors: wanghaijie01
 * @LastEditTime: 2022-02-14 00:17:08
 */

package main

import "net/http"

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}

func main() {
	// 第一种用法
	http.HandleFunc("/hello", helloHandler)
	http.ListenAndServe(":8888", nil)
	// 第二种用法
	mu := http.NewServeMux()
	mu.HandleFunc("/", helloHandler)
	http.ListenAndServe("", mu)

}
