func maximumLength(nums []int, k int) int {
    fs := make(map[int][]int)
    mx := make([]int, k + 2)
    for _, x := range nums {
        if _, ok := fs[x]; !ok {
            fs[x] = make([]int, k + 1)
        }

        f := fs[x]
        for j := k; j >= 0; j-- {
            f[j] = max(f[j], mx[j]) + 1
            mx[j+1] = max(mx[j+1], f[j])
        }
    }

    return mx[k+1]
}