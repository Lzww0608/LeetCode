func findPrefixScore(nums []int) []int64 {
    n := len(nums)
    ans := make([]int64, n)
    mx, pre := 0, 0
    for i, x := range nums {
        mx = max(mx, x)
        ans[i] = int64(pre + mx + x) 
        pre = int(ans[i])
    }

    return ans
}