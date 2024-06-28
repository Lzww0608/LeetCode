func shortestSubarray(nums []int, k int) int {
    n, ans := len(nums), len(nums) + 1
    prefix := make([]int, n + 1)
    for i, x := range nums {
        prefix[i+1] = prefix[i] + x
    }
    dq := []int{}
    for i := 0; i <= n; i++ {
        for len(dq) > 0 && prefix[i] - prefix[dq[0]] >= k {
            ans = min(ans, i - dq[0])
            dq = dq[1:]
        }
        for len(dq) > 0 && prefix[dq[len(dq)-1]] > prefix[i] {
            dq = dq[:len(dq)-1]
        }
        dq = append(dq, i)
    }
    if ans == n + 1 {
        return -1
    }
    return ans
}
