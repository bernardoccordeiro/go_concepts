package main

import "fmt"

func add_f32(x float32, y float32) float32 {
	return x + y
}

func add_f64(x, y float64) float64 {
	return x + y
}

func multiple(a, b string) (string, string) {
	return a, b
}

func pointer_example() {
	x := 15
	a := &x

	fmt.Println(a, *a)
	*a = *a * *a
	fmt.Println(*a, x)
}

func main() {
	var num1, num2 float32 = 5.6, 9.5

	fmt.Println(add_f32(num1, num2)) //works no problem!
	//fmt.Println(add_f64(num1, num2)) //doesn't work since it doesn't cast

	w1, w2 := "Hey", "there"

	fmt.Println(multiple(w1, w2))

	pointer_example()
}
