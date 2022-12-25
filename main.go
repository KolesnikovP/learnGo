package main

import "fmt"

func main() {
	const (
		monday = iota
		tuesday
		wednesday
		thursday
		friday
		saturday
		sunday
	)

	const (
		A int = 45
		B
		C float32 = 3.3
		D
	)

	fmt.Println(monday, tuesday, wednesday, thursday, friday, saturday, sunday) // 0 1 2 3 4 5 6

	fmt.Println(A, B, C, D) // Выведет: 45 45 3.3 3.3
}

/*
В объявлении константы предварительно объявленный идентификатор iota представляет последовательные не типизированные целочисленные константы.
Его значение является индексом соответствующего ConstSpec в объявлении константы, начиная с нуля.
Поскольку он может использоваться в выражениях, он обеспечивает общность, выходящую за рамки простых перечислений.
*/
