func groupAnagrams(strs []string) [][]string {
    groups := make(map[string][]string)
    for _,s := range strs {
        key := SortSlice(s)
        groups[key]= append(groups[key],s)
    }
    result := make([][]string,0)
    for _,group := range groups {
        result = append(result,group)
    }
    return result
}
func SortSlice(s string) string {
   charS := []rune(s)
    sort.Slice(charS, func(i,j int) bool{
        return charS[i] < charS[j]
    })
    return string(charS)
}