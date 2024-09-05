package main

// 从中心向两端扩散的双指针技巧。
func longestPalindrom(s string) string {
	var res string
	for i := 0; i < len(s); i++ {
		// 相当于分奇偶讨论
		s1 := palindrom(s, i, i)
		s2 := palindrom(s, i, i+1)
		if len(res) < len(s1) {
			res = s1
		}
		if len(res) < len(s2) {
			res = s2
		}
	}
	return res
}

func palindrom(s string, l, r int) string {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	return s[l:r]
}
