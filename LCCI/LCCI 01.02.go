const N int = 26
func CheckPermutation(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    cnt := [N]int{}
    for _, c := range s {
        cnt[int(c - 'a')]++
    }
    for _, c := range t {
        cnt[int(c - 'a')]--
    }

    for _, x := range cnt {
        if x != 0 {
            return false
        }
    }
    
    return true
}