func maximumValueSum(nums []int, k int, edges [][]int) int64 {
    n := len(nums)
    a := make([]int, n)
    ans := 0
    for i, x := range nums {
        a[i] = x ^ k - x
        ans += x
    }
    sort.Ints(a)
    for i := n - 1; i >= 1; i -= 2 {
        x := a[i] + a[i-1]
        if x <= 0 {
            break
        } else {
            ans += x 
        }
    }

    return int64(ans)
}