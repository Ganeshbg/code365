package main

import (
	"fmt"
	"strings"
)

func reverseStringList(s []string) []string{
	length := len(s)-1
	if length == 0 {
		return nil
	}

	temp := length
	reverse := make([]string, length+1)
	for temp >= 0 {
		reverse[temp] = s[length-temp]
		temp--
	}

	return reverse
}

func reverseString(s string) string{
	straight := strings.Split(s, ".")
	reverse := reverseStringList(straight)
	if reverse == nil {
		return "Empty string"
	}

	return strings.Join(reverse,".")
}

func main() {
	s := "I.Like.Apples"
	reverse := reverseString(s)
	fmt.Println(reverse)
}

// Given a String S, reverse the string without reversing its individual words. Words are separated by dots.
// https://practice.geeksforgeeks.org/problems/reverse-words-in-a-given-string5459/1