func lexSmallestNegatedPerm(n int, target int64) []int {
    sum := n * (n + 1) / 2
    if target > int64(sum) || (int64(sum) - target) & 1 == 1 || target < int64(-sum) {
        return nil
    }

    d := (sum - int(target)) / 2
    ans := make([]int, n)
    for i := n - 1; i >= 0; i-- {
        if d >= (i + 1) {
            d -= i + 1
            ans[i] = -(i + 1)
        } else {
            ans[i] = i + 1
        }
    }
    sort.Ints(ans)

    return ans
}