func maxProduct(nums []int) int64 {
    mx := slices.Max(nums)
    n := len(nums)
    d := bits.Len(uint(mx))

    if n * n <= n + mx * d {
        ans := 0
        for i := range n {
            for j := range n {
                if nums[i] & nums[j] == 0 {
                    ans = max(ans, nums[i] * nums[j])
                }
            }
        }

        return int64(ans)
    }

    m := 1 << d 
    f := make([]int, m)
    for _, x := range nums {
        f[x] = x
    }

    for i := range m {
        for j := i; j > 0; j -= j & (-j) {
            f[i] = max(f[i], f[i ^ (j & (-j))])
        }
    }

    ans := 0
    for _, x := range nums {
        ans = max(ans, x * f[(m - 1) ^ x])
    }

    return int64(ans)
}