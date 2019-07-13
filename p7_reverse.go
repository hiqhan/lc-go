package leetcode

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

// Reverse a given number
// func Reverse(x int) int {
// 	return reverse(x)
// }
