func canDivideIntoSubsequences(nums []int, k int) bool {
    cnt := make(map[int]int)
    mx := 0 
    for _, x := range nums {
        cnt[x]++
        mx = max(mx, cnt[x])
    }

    return mx * k <= len(nums)
}