package main

import (
	"fmt"
	"strings"
)

/*
simple solution: count max length at every character by iterating the given string
*/
// func lengthOfLongestSubstring(s string) int {
// 	count := make([]int, len(s))
// 	begin := 0
// 	for i, c := range s {
// 		preIndex := strings.LastIndex(s[:i], string(c))
// 		switch preIndex {
// 		case -1:
// 			break
// 		default:
// 			if begin < preIndex+1 {
// 				begin = preIndex + 1
// 			}
// 		}
// 		for j := begin; j <= i; j++ {
// 			count[j]++
// 		}
// 	}

// 	max := 0
// 	for _, c := range count {
// 		if max < c {
// 			max = c
// 		}
// 	}
// 	return max
// }

func lengthOfLongestSubstring(s string) int {
	begin, max := 0, 0
	for i, c := range s {
		preIndex := strings.LastIndex(s[:i], string(c))
		switch preIndex {
		case -1:
			break
		default:
			if offset := i - begin; max < offset {
				max = offset
			}
			if begin < preIndex+1 {
				begin = preIndex + 1
			}
		}
	}
	if offset := len(s) - begin; max < offset {
		max = offset
	}
	return max
}

func main() {
	data := []string{"museuwzbczdejna", "bbb", "abcabcd"}
	for _, s := range data {
		fmt.Println(lengthOfLongestSubstring(s))
	}
}
