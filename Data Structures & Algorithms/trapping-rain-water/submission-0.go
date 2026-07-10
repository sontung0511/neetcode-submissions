func trap(height []int) int {
	left,right := 0, len(height) - 1
	maxLeft:= 0
	maxRight := 0
	total := 0
	for left < right {
		if height[left] < height[right]{
			if maxLeft < height[left]{
				maxLeft = height[left]
			}else{
				total+=maxLeft - height[left]
			}
			left++
		}else{
			if height[right] > maxRight {
				maxRight = height[right]
			}else {
				total += maxRight - height[right]
			}
			right--
		}
	}
	return total
	
}
