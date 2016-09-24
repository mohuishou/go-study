package main

func main() {
	q4_1()
}

func q4_1() {
	sum := 0.0
	x := make([]float64, 10)
	count := len(x)
	for i := 0; i < count; i++ {
		sum += x[i]
	}
	// avg := sum / float64(count)
	println(count)
}
