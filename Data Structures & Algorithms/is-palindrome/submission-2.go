func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	arrayS := []rune(s)
	left,right :=0, len(arrayS)-1

	for left < right {
		for left < right && !unicode.IsDigit(arrayS[left]) && !unicode.IsLetter(arrayS[left]){
			left++
		}
		for left < right && !unicode.IsDigit(arrayS[right]) && !unicode.IsLetter(arrayS[right]){
			right--
		}
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
