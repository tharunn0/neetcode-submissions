func groupAnagrams(strs []string) [][]string {

    res := make([][]string,0,len(strs))
    m := make(map[[26]int][]string)

    for _, str  := range strs {

        var count [26]int

        for _, c := range str {
            count[c-'a']++
        }

        m[count] = append(m[count],str)

    }

    for _, v := range m {
        res = append(res, v)
    }

    

    return res
}
