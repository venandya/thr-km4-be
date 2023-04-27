package main

import "fmt"

func AddElement(array []int, element int, position string) []int {
	if position == "up" {
		return append([]int{element}, array...)
	} else if position == "down" {
		return append(array, element)
	}
	return array
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
