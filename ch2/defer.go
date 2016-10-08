//延迟代码
package main

import (
	"fmt"
)

func main() {
	x := ch2Defer2()
	fmt.Println(x)
}

func ch2Defer() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
	//延迟的函数是按照后进先出(LIFO)的顺序执行,所以上面的代码打印: 4 3 2 1 0 。
	//利用 defer 甚至可以修改返回值
}

//利用defer可以改变返回值
func ch2Defer2() (x int) {
	defer func() {
		x++
	}()
	return 0 //最后的返回结果为1
}
