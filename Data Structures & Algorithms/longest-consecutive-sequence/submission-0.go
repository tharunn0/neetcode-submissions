func longestConsecutive(nums []int) int {
	m := make(map[int]bool)
	for _, n := range nums {
		m[n] = true
	}
	lcq, mx := 1, 0
	for _, n := range nums {

		if _, ok := m[n-1]; ok {
			continue
		} 

		nextn := n +1
		for  {

			if _, ok := m[nextn]; ok {
				lcq++
				nextn++
			} else {
				break
			}
		}
		mx = max(mx, lcq)
		lcq = 1
	}
	return mx
}
