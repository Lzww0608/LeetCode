func minimumRecolors(blocks string, k int) int {
    ans := k
    n := len(blocks)
    cnt := 0
    for l, r := 0, 0; r < n; r++ {
        if blocks[r] == 'B' {
            cnt++
        }
        if r - l + 1 == k {
            ans = min(ans, k - cnt)
            if blocks[l] == 'B' {
                cnt--
            }
            l++
        }
    }

    return ans
}