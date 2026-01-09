func minLength(nums []int, k int) int {
    n := len(nums)
    ans := n + 1

    cnt := make(map[int]int)
    sum := 0
    for l, r := 0, 0; r < n; r++ {
        x := nums[r]
        if cnt[x]++; cnt[x] == 1 {
            sum += x 
        }

        for sum >= k {
            ans = min(ans, r - l + 1)
            y := nums[l]
            if cnt[y]--; cnt[y] == 0 {
                sum -= y
            }
            l++
        }
    }

    if ans > n {
        return -1
    }
    return ans
}