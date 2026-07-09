func twoSum(nums []int, target int) []int {
  var sum []int
    for i := 0; i < len(nums);i++{
      for j := i+1;j<len(nums);j++{
        total := nums[i]+nums[j]
        if total == target {
          sum = append(sum,i,j)
          break
        }
      
      }
    }
  return sum
}
