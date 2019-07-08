package main

import "fmt"

func reverse(x int) int {
	var rev int32
	for x >= 10 || x <= -10 {
		rev = rev*10 + int32(x)%10
		x /= 10
	}

	if rev*10/10 != rev {
		return 0
	} else {
		return int(rev*10 + int32(x)%10)
	}
}

func main() {
	data := []int{123, -123, 1200, 0, 1534236469}
	for _, i := range data {
		fmt.Println(i, reverse(i))
	}
}
