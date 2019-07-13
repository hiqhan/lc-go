package leetcode

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	lenS, cs := len(s), string("")
	for i := 0; i < numRows; i++ {
		for j := i; j < lenS; j += (numRows - 1) * 2 {
			cs += string(s[j])
			if i == 0 || i == numRows-1 {
				continue
			}
			if v := j + (numRows-1-i)*2; v < lenS {
				cs += string(s[v])
			}
		}
	}
	return cs
}

// Convert string to zigzag and return a string read from zigzag
// func Convert(s string, n int) string {
// 	return convert(s, n)
// }
