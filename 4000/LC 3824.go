func minimumK(nums []int) int {
    
    check := func(k int) bool {
        cnt := 0
        for _, x := range nums {
            cnt += (x + k - 1) / k
        }

        return cnt <= k * k
    }

    l, r := 0, 100_000
    for l + 1 < r {
        mid := l + ((r - l) >> 1)
        if check(mid) {
            r = mid
        } else {
            l = mid 
        }
    }

    return r
}