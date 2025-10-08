func resultArray(nums []int, k int) []int64 {
    f := make([]int64, k)
    ans := make([]int64, k)

    for _, x := range nums {
        x %= k
        g := make([]int64, k)
        g[x] = 1
        for y := range k {
            g[x * y % k] += f[y]
        }

        f = g 
        for y := range k {
            ans[y] += f[y]
        }
    }

    return ans
}