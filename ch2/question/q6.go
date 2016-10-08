package main

func main() {
	a, b := swip(1, 2)
	println(a, b)
}

func swip(x, y int) (a, b int) {
	a = x
	b = y
	if x > y {
		a = y
		b = x
	}
	return
}
