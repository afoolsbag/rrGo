package main

import "fmt"
import "rrgo"

// Go 中的 main 函数，既没有参数，也没有返回类型。
func main() {
	// 用于调试环境
	println("hello, world")

	// 用于生产环境
	fmt.Println("hello, world")

	fmt.Println(rrgo.NextBirthdayAnniversaryCountdown().Hours() / 24)
}
