package main

//流程控制语句
func main() {
	ch1If()
	ch1For()
	ch1Range()
	ch1Switch()
}

//if语句
func ch1If() {
	//注意：if的大括号必须
	x := -10
	if x > 0 {
		print(x)
	} else {
		print("x是负数")
	}
}

//for语句,有三种形式
func ch1For() {
	//常规for循环形式
	for i := 0; i < 10; i++ {
		print(i)
	}

	//while模式
	j := 2
	for j < 5 {
		print(j)
		j++
	}

	// 死循环
	for {
		print(j)
		break
	}

	//嵌套循环终止指定标签的循环层次
}

func ch1Range() {
	//感觉有点类似php当中的foreach
	//用于slice、 array、 string、 map 和 channel
	list := []string{"a", "b", "c"}
	for k, v := range list {
		println(k)
		println(v)
	}

}

func ch1Switch() {
	//两个注意的点
	//1.在go当中switch可以没有表达式，在没有表达式的时候会匹配true=》可以用于书写if-else结构
	//2.不会匹配成功之后自动往下尝试，想要让其自动向下尝试可以使用fallthrough，相当于每一条case语句都已经自带了break
	//3.switch的条件和case的类型必须相同，default的使用方法一样

	//验证
	a := 1
	switch {
	case a < 0:
		println("a是负数！")
	case a == 0:
		println("a是0")
	case a > 0:
		println("a是正数")
	}

	b := 10
	switch b {
	case 0:
		println("b是0")
	case 10:
		println("b是10")
	}

}
