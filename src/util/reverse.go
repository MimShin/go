package util

/*
func Reverse(s string) string {
	if len(s) < 2 {
		return s
	}
	l := len(s) - 1
	return s[l:] + Reverse(s[:l])
}
*/

func Reverse(s string) string {
	r := []rune(s)
	
	for i, j := 0, len(s)-1; i<j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
