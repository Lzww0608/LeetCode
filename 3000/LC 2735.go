func minCost(nums []int, x int) (ans int64) {
    for _, x := range nums {
        ans += int64(x)
    }

    n := len(nums)
    a := append([]int(nil), nums...)
    for i := range nums {
        var s int64
        for j := range nums {
            a[j] = min(a[j], nums[(j - i + n) % n])
            s += int64(a[j])
        }
        ans = min(ans, s + int64(x) * int64(i))
    }

    return 
}