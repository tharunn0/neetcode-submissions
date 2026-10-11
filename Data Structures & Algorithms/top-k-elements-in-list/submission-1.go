// type entry struct {
// 	value int
// 	freq int
// }

func topKFrequent(nums []int, k int) []int {
// METHOD 1 - MinHeap with len K
// create map of key=num , value freq

// convert it to entry slice

// convert it to a min heap based on the freq value field and popping the min values

// pop k times and return the value

// METHOD - 2 frequency map with bucket sort
// create map of key=num, value=freq
m := make(map[int]int)
for _, n := range nums {
	m[n]++
}

// create a slice of len(nums) as max freq will be <= n
arr := make([][]int, len(nums)+1)

for k, v := range m {
	arr[v] = append(arr[v], k)
} 
res := []int{}
count := 0
for i := len(arr)-1; i > 0; i-- {
	ar := arr[i]
	for j := len(ar)-1; j >= 0; j-- {
		if count == k {
		return res
		}

		res = append(res, ar[j])

		count++
	}
}
return res
}
