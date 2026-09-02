func isAnagram(s string, t string) bool {

    if len(s) != len(t) {
        return false
    }
    
   var arr [26]rune

   for i := range s {
        arr[s[i]-'a']++
        arr[t[i]-'a']--
   }

   for _, v := range arr {
    if v != 0 {
        return false
    }
   }
return true
}
