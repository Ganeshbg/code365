package main

import (
	"fmt"
)

func linearSearch(list []int, key int) bool {
	for _, v := range list {
		if v == key {
			return true
		}
	}
	return false
}

func main() {
	list := []int{1, 2, 3, 4, 5, 6, 78, 93, 98, 34, 344}
	key := 108
	found := linearSearch(list, key)
	if found {
		fmt.Println("Key Found")
	} else {
		fmt.Println("Key Not Found")
	}

}
