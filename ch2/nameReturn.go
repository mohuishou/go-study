//命名返回值
package main

import "os"

func main() {
	y, err := test2("这是一个测试字符串")
	if err != nil {
		println("发生了一个错误！")
	}
	println(y)
}

func test2(str string) (y int, err error) {
	file, err := os.Create("test.txt")
	if err != nil {
		return
	}
	y, err = file.WriteString(str)
	return
}
