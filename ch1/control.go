package main

//流程控制语句
func main() {
	ch1If()
	ch1For()
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
}

func ch1Switch() {

}
