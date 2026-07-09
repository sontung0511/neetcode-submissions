func isAnagram(s string, t string) bool {
   if len(s) != len(t) {
        return false
    }
  arrT := strings.Split(t,"")
  arrS := strings.Split(s,"")
  sort.Strings(arrT)
  sort.Strings(arrS)
  a := true
  for i,value := range arrT {
    if arrS[i] != value {
      a = false
      break
    }
  }
  return a
}
