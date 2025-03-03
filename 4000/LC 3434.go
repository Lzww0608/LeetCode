const N int = 55
func maxFrequency(nums []int, k int) (ans int) {
    f := make([]int, N)
    cur := 0
    for _, x := range nums {
        if x == k {
            cur++
            ans++
        } else {
            f[x] = max(f[x], cur) + 1 
            ans = max(ans, f[x])
        }
    }

    return
}