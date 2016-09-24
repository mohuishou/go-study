/*
	编写一个函数用于计算一个 float64 类型的 slice 的平均值。
*/
package main

func main() {
	// x := make([]float64, 3)
	// x[0] = 1.3
	// x[1] = 1.4
	// x[2] = 1.5
	x := []float64{1.2, 1.3, 1.4}
	y := avg(x)
	println(y)
}

func avg(x []float64) (y float64) {
	sum := 0.0
	if len(x) == 0 {
		return 0.0
	}
	for _, v := range x {
		sum += v
	}
	y = sum / float64(len(x))
	return
}
