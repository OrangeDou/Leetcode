package main

/*
回文子串必有一个中心，该中心可能是对称轴，也可能是某个char，从它向两边扩散可以找到以对称轴为中心的子串，循环所有元素，每个char都可能是对称轴
*/
func longestPalindrome(s string) string {
	var res string
	for i := 0; i < len(s); i++ {
		s1 := find(s, i, i)
		s2 := find(s, i, i+1)
		if len(res) < len(s1) {
			res = s1
		}

		if len(res) < len(s2) {
			res = s2
		}
	}
	return res
}

func find(s string, l, r int) string {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	return s[l+1 : r]
}
