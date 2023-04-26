package main

import "fmt"

func howManyElements(data []any) int {
	return len(data)
}

func main() {
	input1 := []any{1, 2, 3, 4, 5}
	fmt.Println(howManyElements(input1)) // Output: 5

	input2 := []any{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(howManyElements(input2)) // Output: 10

	input3 := []any{1, 1, 1, 5, 5, 5}
	fmt.Println(howManyElements(input3)) // Output: 6

	input4 := []any{"Edo", "Budi", "Joko", "Tono"}
	fmt.Println(howManyElements(input4)) // Output: 4

	input5 := []any{"Edo", "Budi", "Joko", "Tono", "Edo", "Budi", "Joko", "Tono"}
	fmt.Println(howManyElements(input5)) // Output: 8

	input6 := []any{true, false, true, false, true, false}
	fmt.Println(howManyElements(input6))
}
