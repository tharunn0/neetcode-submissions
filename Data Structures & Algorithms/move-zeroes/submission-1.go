func moveZeroes(nums []int) {
	zi, i := 0 , 0

	for i < len(nums) {
		if nums[i] != 0 {
			nums[zi], nums[i] = nums[i], nums[zi]
		i++
		zi++
		} else {
			i++
		}
	}
}
