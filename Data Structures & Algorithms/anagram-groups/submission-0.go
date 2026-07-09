func groupAnagrams(strs []string) [][]string {
  res := make(map[string][]string)
  for _,s := range strs {
    sorts := SortRune(s)
    res[sorts] = append(res[sorts],s)
  }
  var myMatrix [][]string
  for _,value := range res {
    myMatrix = append(myMatrix,value)
  }

  return myMatrix
}
func SortRune(s string)string{
  characters := []rune(s)
  sort.Slice(characters, func(i, j int) bool {
    return characters[i] < characters[j]
  })
  return string(characters)
}