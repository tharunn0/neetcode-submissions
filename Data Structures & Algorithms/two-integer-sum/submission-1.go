func twoSum(nums []int, target int) []int {

    // var i,j int 
    // j = len(nums)-1
    // for i < j  {
    //     if nums[i]+nums[j] == target {
    //         return []int{i,j}
    //     } else if nums[i]+nums[j] > target {
    //         j--
    //     } else {
    //         i++
    //     }
    // }
    // return []int{}

    m := make(map[int]int)
    
    for i,v := range nums {
        n := target-v
        if idx,ok := m[n];ok {
            return []int{idx,i}
        }
        m[v] = i
    }
    return []int{}
}
