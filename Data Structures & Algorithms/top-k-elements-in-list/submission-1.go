type Pair struct {
	Key   int
	Value int
}
func topKFrequent(nums []int, k int) []int {
  var dublicate []int
  mapA := make(map[int]int)
  for _, num := range nums {
      mapA[num]++
  }
  var pairs []Pair
  for k,v := range mapA{
    maxValue := 0
    if v > maxValue {
      pairs = append(pairs,Pair{Key:k,Value:v})
    }
  }
  sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Value > pairs[j].Value
	})
  var topN []Pair
	if len(pairs) < k {
		topN = pairs
	} else {
		topN = pairs[:k]
	}
	for _, value := range topN {
		dublicate = append(dublicate,value.Key)
	}
  return dublicate
}
