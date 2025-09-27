func getLargestOutlier(nums []int) int {
    sum := 0
    cnt := make(map[int]int)
    for _, x := range nums {
        sum += x
        cnt[x]++
    }

    ans := math.MinInt
    for _, x := range nums {
        cur := sum - x 
        cnt[x]--
        if cur & 1 == 0 && cnt[cur / 2] > 0 {
            ans = max(ans, x)
        }
        cnt[x]++
    }

    return ans
}