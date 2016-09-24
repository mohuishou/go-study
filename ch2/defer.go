//延迟代码
package main

import (
	"fmt"
)

func main() {
	ch2Defer()
}

func ch2Defer() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
	//延迟的函数是按照后进先出(LIFO)的顺序执行,所以上面的代码打印: 4 3 2 1 0 。
	//利用 defer 甚至可以修改返回值
}
