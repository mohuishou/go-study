package main

import (
	"fmt"
)

/**
 *变量、类型
 */
func main() {
	ch1Int()
	ch1Str()
}

/**
 * int变量
 */
func ch1Int() {
	//变量声明
	// note: 变量声明未使用会报错
	// var a int
	// a = 10 //仅仅赋值不算使用
	// b := 11

	//成组声明，const和impor也可以这样使用
	// var (
	// 	c int
	// 	d float32
	// )

	// c = 13
	// d = 11

	// fmt.Println(c + d) //error 不同类型的不能直接相加

	//下划线是特殊变量，赋值给他的值将会被丢弃
	//平行赋值
	_, f := 1, 5
	fmt.Println(f)
}

/**
 * 字符串变量
 * 字符串变量需要注意，字符串一旦赋值不能在直接改变
 */
func ch1Str() {
	s := "hello world"

	// 字符串的修改

	//尝试通过php的方式修改
	//s[0] = 1 error 报错，不能直接这样修改

	//修改方式：先转换为数组，修改之后在转换为字符串
	c := []rune(s) //rune是int32的一个别名，可以用于遍历字符串当中的字符
	c[0] = 'a'     //注意：单引号包含的是字符类型，双引号是字符串类型，这里不能使用双引号
	s2 := string(c)

	fmt.Println(s2)

	//多行字符串
	s3 := "abc" +
		"def"
	fmt.Println(s3)
}
