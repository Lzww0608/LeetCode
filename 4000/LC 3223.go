const N int = 26
func minimumLength(s string) (ans int) {
    n := len(s)
    if n < 3 {
        return n
    }

    cnt := [N]int{}
    for i := range s {
        cnt[int(s[i] - 'a')]++
    }

    for _, x := range cnt {
        if x & 1 == 1 {
            ans += min(x, 1)
        } else {
            ans += min(x, 2)
        }
    }

    return
}