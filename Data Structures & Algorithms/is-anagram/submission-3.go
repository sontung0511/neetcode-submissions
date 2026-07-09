func isAnagram(s string, t string) bool {
	arrayS := []rune(s)
	arrayT := []rune(t)
	if len(arrayS) != len(arrayT){
		return false
	}
	count := make(map[rune]int)
	for _,sRune := range arrayS {
		count[sRune]++
	}
	for _,tRune := range arrayT {
		count[tRune]--
		if count[tRune] < 0{
			return false
		}
	}
	return true
}
