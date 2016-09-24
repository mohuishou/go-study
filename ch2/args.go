//变参
package main

func main() {
	ch2Args1(1, 5, 6)
}

//参数是int类型的一个slice
func ch2Args1(arg ...int) {
	for _, v := range arg {
		println(v)
	}
}

//如果不指定变参的类型,默认是空的接口 interface{}
