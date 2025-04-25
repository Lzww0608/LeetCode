func countInterestingSubarrays(nums []int, modulo int, k int) int64 {
    ans := 0
    n := len(nums)
    pre := make([]int, n + 1)
    for i, x := range nums {
        if x % modulo == k {
            pre[i+1] = pre[i] + 1 
        } else {
            pre[i+1] = pre[i]
        }
    }

    // (pre[r] - pre[l]) % modulo == k % modulo
    // (pre[r] - k) % modulo == pre[l] % modulo 
    cnt := make(map[int]int)
    for _, x := range pre {
        ans += cnt[(x - k + modulo) % modulo]
        cnt[x % modulo]++
    }

    return int64(ans)
}