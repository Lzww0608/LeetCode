func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    cnt := make([]int, 26)
    f := true

    for i := range s {
        if s[i] != t[i] {
            f = false
        }

        cnt[int(s[i] - 'a')]++
        cnt[int(t[i] - 'a')]--
    }

    for _, x := range cnt {
        if x != 0 {
            return false
        }
    }

    return !f
}