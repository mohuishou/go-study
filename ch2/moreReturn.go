//多值返回
package main

func main() {
	y, err := test(10)
	if err != nil {
		println("有一个错误")
	}
	println(y)
}

func test(x int) (y int, err error) {

	y = x + 1
	return y, err
}
