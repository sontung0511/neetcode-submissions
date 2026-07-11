func lengthOfLongestSubstring(s string) int {
	seen := make(map[byte]bool)
	left := 0
	maxLen := 0
	for right := 0;right < len(s);right++{
		for seen[s[right]] {
			delete(seen,s[left])
			left++
		}
		seen[s[right]] = true
		lenCal := right - left + 1
		if maxLen < lenCal {
			maxLen = lenCal
		}
	}
	return maxLen
}
