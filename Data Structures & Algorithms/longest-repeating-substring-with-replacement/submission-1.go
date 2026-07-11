func characterReplacement(s string, k int) int {
	count := make(map[byte]int)
	left := 0
	maxFreq := 0
	result := 0
	for right := 0;right < len(s);right++{
		ch := s[right]
		count[ch]++
		if count[ch] > maxFreq {
			maxFreq = count[ch]
		}
		winLen := right - left + 1
		for winLen - maxFreq > k{
			count[s[left]]--
			left++
			winLen = right - left + 1
		}
		if winLen > result {
			result = winLen
		}
	}
	return result
}
