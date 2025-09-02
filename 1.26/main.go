package main

import (
	"fmt"
	"strings"
)

func deleteDuplicate(str string) bool {
	str = strings.ToLower(str)

	for i := 0; i < len(str); i++ {
		for j := i + 1; j < len(str); j++ {
			if str[i] == str[j] {
				return false
			}
		}
	}

	return true
}

func main() {
	s := []string{"abcd", "abCdefAF", "aabcd"}

	for i := range s {
		fmt.Println(deleteDuplicate(s[i]))
	}
}
