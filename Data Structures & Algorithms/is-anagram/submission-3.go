func isAnagram(s string, t string) bool {

    if len(s) != len(t) {
        return false
    }

    freq := make([]int, 26)

    for i := range s {
        freq[rune(s[i])-'a']++
        freq[rune(t[i])-'a']--
    }

    for _, n := range freq {
        if n != 0 {return false}
    }
    return true
}
