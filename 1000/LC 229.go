func majorityElement(nums []int) (ans []int) {
    a, b := math.MinInt32, math.MinInt32 
    ca, cb := 0, 0 

    for _, x := range nums {
        if x == a {
            ca++
        } else if x == b {
            cb++
        } else if ca == 0 {
            a, ca = x, 1
        } else if cb == 0 {
            b, cb = x, 1
        } else {
            ca, cb = ca - 1, cb - 1
        }
    }

    ca, cb = 0, 0 
    for _, x := range nums {
        if x == a {
            ca++
        } else if x == b {
            cb++
        }
    }

    if ca > len(nums) / 3 {
        ans = append(ans, a)
    }
    if cb > len(nums) / 3 {
        ans = append(ans, b)
    }

    return
}