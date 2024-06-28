func longestEqualSubarray(nums []int, k int) (ans int) {
    m := make([][]int, len(nums) + 1)
    for i, x := range nums {
        m[x] = append(m[x], i)
    }
    for _, v := range m {
        n := len(v)
        if n <= ans {
            continue
        }
        l, r := 0, 0
        for r < n {
            for v[r] - v[l] - (r - l) > k {
                l++
            }
            ans = max(ans, r - l + 1)
            r++
        }
    }

    return 
}
