func mostFrequentEven(nums []int) int {
    cnt := make(map[int]int)
    ans, mx := -1, 0
    for _, x := range nums {
        if x & 1 == 0 {
            cnt[x]++
            if cnt[x] > mx || cnt[x] == mx && x < ans {
                ans, mx = x, cnt[x]
            } 
        }
    }

    return ans
}