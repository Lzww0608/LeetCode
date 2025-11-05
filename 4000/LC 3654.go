func minArraySum(nums []int, k int) int64 {
    n := len(nums)
    nxt := make([]int, n)
    pre := make(map[int]int)
    pre[0] = 0

    sum := 0
    for i, x := range nums {
        nxt[i] = n + 1
        sum += x 

        t := sum % k
        if v, ok := pre[t]; ok {
            nxt[v] = i + 1 
        } 
        pre[t] = i + 1
    }

    f := make([]int, n + 1)
    for i := n - 1; i >= 0; i-- {
        f[i] = f[i + 1] + nums[i]
        if nxt[i] != n + 1 {
            f[i] = min(f[i], f[nxt[i]])
        }
    }

    return int64(f[0])
}