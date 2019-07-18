package leetcode

func myAtoi(str string) int {
	if len(str) == 0 {
		return 0
	}

	var sign, n int // 1: positive, -1: negative
	leading := true
	for _, c := range str {
		if c == ' ' && leading {
			continue
		}
		if c == '+' && leading {
			sign, leading = 1, false
			continue
		}
		if c == '-' && leading {
			sign, leading = -1, false
			continue
		}
		if 48 <= c && c <= 57 {
			if sign >= 0 {
				n = n*10 + int(c-48)
				if n > (1<<31 - 1) {
					return 1<<31 - 1
				}
			} else {
				n = n*10 - int(c-48)
				if n < -(1 << 31) {
					return -(1 << 31)
				}
			}
			leading = false
		} else {
			break
		}
	}

	return n
}
