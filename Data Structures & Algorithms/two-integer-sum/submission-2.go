func twoSum(nums []int, target int) []int {
    seenNum := make(map[int]int)
    for i,num := range nums {
        needNum := target - num
        if idx,ok := seenNum[needNum];ok{
            return []int{idx,i}
        }
        seenNum[num]=i
    }
    return nil
}
