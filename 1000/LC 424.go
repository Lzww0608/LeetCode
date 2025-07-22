const N int = 26
func characterReplacement(s string, k int) int {
    ans := k 
    n := len(s)
    for i := range N {
        cnt := 0
        for l, r := 0, 0; r < n; r++ {
            if int(s[r] - 'A') == i {
                cnt++
            }

            for r - l + 1 > cnt + k {
                if int(s[l] - 'A') == i {
                    cnt--
                }
                l++
            }

            ans = max(ans, r - l + 1)
        }
    }

    return ans
}