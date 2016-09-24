package main

func main() {
	ch1Map()
}

func ch1Array() {
	//数组初始化
	// var arr [10]int
	// b := [...]int{123, 1, 2}
	//多位数组初始化
	// c := [2][2]int{{123, 123}, {23, 32}}
}

func ch1Slice() {
	//和数组类似但是长度可以改变
	//slice是引用类型
	// sl := make([]int, 10) //引用类型通过make创建
}

func ch1Map() {
	monthdays := map[string]int{
		"Jan": 31, "Feb": 28, "Mar": 31,
		"Apr": 30, "May": 31, "Jun": 30,
		"Jul": 31, "Aug": 31, "Sep": 30,
		"Oct": 31, "Nov": 30, "Dec": 31, //逗号是必须的
	}
	println(monthdays["Oct"])

	//增加元素
	monthdays["Undecim"] = 30

	//检查元素是否存在
	var (
		v int
	)
	v, ok := monthdays["Jan"]

	println(ok, v)

	//删除元素
	delete(monthdays, "Apr")
}
