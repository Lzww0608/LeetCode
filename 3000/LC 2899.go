func lastVisitedIntegers(nums []int) (ans []int) {
    seen := []int{}
    
    cnt := 0
    for _, x := range nums {
        if x > 0 {
            cnt = 0
            seen = append(seen, x)
        } else {
            cnt++
            n := len(seen)
            if cnt > n {
                ans = append(ans, -1)
            } else {
                ans = append(ans, seen[n - cnt])
            }
        }
    }

    return
}