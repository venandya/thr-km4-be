package main

import "fmt"

func AddElement(array []int, element int, position string) []int {
	if position == "up" {
		newArray := make([]int, len(array)+1)
		newArray[0] = element
		copy(newArray[1:], array)
		return newArray
	} else if position == "down" {
		return append(array, element)
	} else {
		return array
	}
}

func main() {
	array := []int{1, 2, 3, 4, 5}
	element := 6
	position := "up"

	newArray := AddElement(array, element, position)
	fmt.Println(newArray)

	position = "down"
	newArray = AddElement(array, element, position)
	fmt.Println(newArray)
}
