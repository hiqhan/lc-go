package leetcode

// LongestPalindrome find the longest palindrome substring
// func LongestPalindrome(s string) string {
// 	return longestPalindrome(s)
// }

func longestPalindrome(s string) string {
	result := make([][]bool, len(s))
	for i := range result {
		result[i] = make([]bool, len(s))
	}

	var palindrome string
	for distance := 0; distance < len(s); distance++ {
		for i := 0; i < len(s)-distance; i++ {
			j := i + distance
			switch distance {
			case 0:
				result[i][j] = true
			case 1:
				result[i][j] = result[i][i] && s[i] == s[j]
			default:
				result[i][j] = result[i+1][j-1] && s[i] == s[j]
			}

			if result[i][j] {
				palindrome = s[i : j+1]
			}
		}
	}

	// for _, v := range result {
	// 	fmt.Println(v)
	// }
	// for distance := len(s) - 1; distance >= 0; distance-- {
	// 	for i := 0; i < len(s)-distance; i++ {
	// 		j := i + distance
	// 		if result[i][j] == true {
	// 			return s[i : j+1]
	// 		}
	// 	}
	// }

	return palindrome
}
