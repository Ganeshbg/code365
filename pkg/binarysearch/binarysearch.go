package main

import (
	"fmt"
	"insertionsort"
)

func binarySearch(list []int, key int)  bool{
	length := len(list)
	low := 0
	mid := length/2
	high := length-1
	for low <= high {
		mid = (low + high)/2
		if list[mid] > key {
			high = mid -1
		} else if list[mid] < key {
			low = mid + 1
		} else {
			return true
		}
	}
	return false
}

func main() {
	list := []int{8,3,4,55,2,9,32}
	list = insertionsort.InsertionSort(list)
	fmt.Println(list)
	found := binarySearch(list, 8)
	if found {
		fmt.Println("Item Found")
	} else {
		fmt.Println("Item not found")
	}
}