const N int = 26
func validSubstringCount(s string, t string) int64 {
    ans := 0
    cnt := make([]int, N)
    for i := range t {
        x := int(t[i] - 'a')
        cnt[x]++
    }

    n := len(s)
    dif := 0
    for _, x := range cnt {
        if x > 0 {
            dif++
        }
    }

    for l, r := 0, 0; r < n; r++ {
        x := int(s[r] - 'a')
        if cnt[x]--; cnt[x] == 0 {
            dif--
        }

        for dif == 0 {
            ans += n - r
            y := int(s[l] - 'a')
            if cnt[y]++; cnt[y] == 1 {
                dif++
            }
            l++
        }
    }

    return int64(ans)
}