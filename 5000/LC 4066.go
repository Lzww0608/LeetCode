func maxEqualAdjacentPairs(nums []int) (ans int) {
    n := len(nums)
    cnt := make(map[[2]int]int)
    base, mx := 0, 0
    for i := 1; i < n; i++ {
        x, y := nums[i - 1], nums[i]
        if x > y {
            x, y = y, x
        } 
        if x == y {
            base++
        } else {
            cnt[[2]int{x, y}]++
            mx = max(mx, cnt[[2]int{x, y}])
        }
        
        ans = max(ans, base + mx)
    }

    return
}