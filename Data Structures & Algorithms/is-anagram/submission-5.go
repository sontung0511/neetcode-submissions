func isAnagram(s string, t string) bool {
    arrayS := []rune(s)
    arrayT := []rune(t)
    if len(arrayT) != len(arrayS){
        return false
    }
    countSeen := make(map[rune]int)
    for _,numS := range arrayS {
        countSeen[numS]++
    }
    for _,numT := range arrayT {
        countSeen[numT]--
        if countSeen[numT] < 0 {
            return false
        }
    }
    return true
}
