package main

import "fmt"

func removeLeftRight(arr []any, position string) []any {
	if position == "left" {
		return arr[1:]
	} else {
		return arr[:len(arr)-1]
	}
}

func main() {
	array1 := []any{1, 2, 3, 4, 5}
	fmt.Println(removeLeftRight(array1, "left"))  // [2 3 4 5]
	fmt.Println(removeLeftRight(array1, "right")) // [1 2 3 4]

	array2 := []any{"Edo", "Budi", "Joko", "Tono"}
	fmt.Println(removeLeftRight(array2, "left"))  // [Budi Joko Tono]
	fmt.Println(removeLeftRight(array2, "right")) // [Edo Budi Joko]
}
