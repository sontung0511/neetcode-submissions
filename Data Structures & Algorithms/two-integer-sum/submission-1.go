func twoSum(nums []int, target int) []int {
	seen := make(map[int]int,len(nums))
	for i,num := range nums {
		numNeed := target - num
		if idx,ok := seen[numNeed];ok{
			return []int{idx,i}
		}
		seen[num]=i
	}
	return nil
}
