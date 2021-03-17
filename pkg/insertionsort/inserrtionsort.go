package main

import (
	"fmt"
)

func insertionSort(list []int) []int{
	for i := 1; i< len(list); i++ {
		j := i-1
		temp := list[i]
		for j>=0 && list[j] > temp {
			list[j+1] = list[j]
			j--
		}
		list[j+1] = temp
	}
	return list
}

func main() {
	list := []int{4,5,2,7,8,33,57,3,2,8,98}
	fmt.Println(insertionSort(list))
}