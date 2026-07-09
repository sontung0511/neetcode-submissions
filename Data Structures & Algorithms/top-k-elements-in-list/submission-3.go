func topKFrequent(nums []int, k int) []int {
    countNum := make(map[int]int)
    for _,num := range nums {
        countNum[num]++
    }
    type group struct {
        key int
        value int
    }
    arrayCount := make([]group, 0)
    for count,num := range countNum {
        arrayCount = append(arrayCount,group{key:count,value:num})
    }
    sort.Slice(arrayCount, func(i,j int) bool {
        return arrayCount[i].value > arrayCount[j].value
    })
    result := make([]int,0)
    for i := 0;i < k;i++ {
        result = append(result,arrayCount[i].key)
    }
    return result
}
