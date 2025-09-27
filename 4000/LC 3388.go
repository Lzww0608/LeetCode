func Z(s []int) []int {
    n := len(s)
    z := make([]int, n)
    for i, l, r := 1, 0, 0; i < n; i++ {
        z[i] = max(min(z[i - l], r - i + 1), 0)
        for i + z[i] < n && s[i + z[i]] == s[z[i]] {
            l, r = i, i + z[i]
            z[i]++
        }
    }

    return z
}

func beautifulSplits(nums []int) (ans int) {
    n := len(nums)
    z := Z(nums)
    for i := 1; i < n - 1; i++ {
        f := Z(nums[i:])
        for j := i + 1; j < n; j++ {
            if i + i <= j && z[i] >= i || f[j - i] >= j - i {
                ans++
            } 
        }
    }

    return
}