import "slices"
func threeSum(nums []int) [][]int {

	slices.Sort(nums)
	res := [][]int{}

	a, b, c := 0, 1, len(nums)-1

	for a < len(nums)-2 {

        if a > 0 && nums[a] == nums[a-1] {
            a++
            continue
        }

		b = a+1
		c = len(nums)-1
		for b < c {
			
			sum := nums[a] + nums[b] + nums[c]
			if sum == 0 {
				res = append(res, []int{nums[a], nums[b], nums[c]})
                c--
                b++

                for b < c && nums[b] == nums[b-1] {
                    b++
                }
                for c > b && nums[c] == nums[c+1] {
                    c--
                }
			} else if sum > 0 {
				c--
			} else {
				b++
			}
		}
		a++
	}

	return res
}
