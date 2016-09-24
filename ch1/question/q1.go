package main

import "fmt"

func main() {
	q1_1()

	q1_2()

	q1_3()
}

func q1_1() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func q1_2() {
	i := 0
I:
	fmt.Println(i)
	i++
	if i < 10 {
		goto I
	}
}

func q1_3() {
	a := [3]int{1, 5, 2}
	for i := 0; i < 3; i++ {
		fmt.Println(a[i])
	}
}
