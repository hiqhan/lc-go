package leetcode

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	i, y := x, 0
	for i > 0 {
		y = y*10 + i%10
		i = i / 10
	}
	return x == y
}

// IsPalindrome Judge a given number is a palindrome or not
func IsPalindrome(x int) bool {
	return isPalindrome(x)
}
