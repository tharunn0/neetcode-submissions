func isPalindrome(s string) bool {

	i, j := 0, len(s)-1

	for i < j {
		if isSame(s[i], s[j]) {
			j--
			i++
		} else if !isValidChar(s[i]) {
			i++
		} else if !isValidChar(s[j]) {
			j--
		} else {
			return false
		}
		
	}
return true
}

func isSame(a, b byte) bool {
	return unicode.ToLower(rune(a)) == unicode.ToLower(rune(b))
}

func isValidChar(a byte) bool {
	return (a <= 'Z' && a >= 'A') ||
	(a <= 'z' && a >= 'a') ||
	(a <= '9' && a >= '0') 
}
