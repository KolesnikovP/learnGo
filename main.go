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

	fmt.Println(monday, tuesday, wednesday, thursday, friday, saturday, sunday)
}

/*
В объявлении константы предварительно объявленный идентификатор iota представляет последовательные не типизированные целочисленные константы.
Его значение является индексом соответствующего ConstSpec в объявлении константы, начиная с нуля.
Поскольку он может использоваться в выражениях, он обеспечивает общность, выходящую за рамки простых перечислений.
*/
