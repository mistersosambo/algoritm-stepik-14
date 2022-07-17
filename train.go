package main

import "fmt"

func main() {
	var a int
	sum := 0
	sum1 := 0
	sum2 := 0
	fmt.Scan(&a)
	//По скольку введенное число не превышает 10^7 справедливым будет,что максимум итерация пробега по сумме числа это 3
	for a != 0 {
		r := a % 10
		sum += r
		a /= 10
	}
	for sum != 0 {
		r := sum % 10
		sum1 += r
		sum /= 10
	}
	for sum1 != 0 {
		r := sum1 % 10
		sum2 += r
		sum1 /= 10
	}
	fmt.Println(sum2)
}
