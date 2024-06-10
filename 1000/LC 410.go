func splitArray(nums []int, k int) int {
    sum, mx := 0, nums[0]
    for _, x := range nums {
        sum += x
        mx = max(mx, x)
    }
    if k == 1 {
        return sum
    }

    check := func(target int) int {
        cnt, s := 1, 0
        for _, x := range nums {
            if s + x <= target {
                s += x
            } else {
                s = x
                cnt++
            }
        }
        return cnt
    }


    l, r := mx, sum
    for l < r {
        mid := l + ((r - l) >> 1)
        cnt := check(mid)
        fmt.Println(mid, cnt)
        if cnt > k {
            l = mid + 1
        } else {
            r = mid 
        }
    }

    return l
}
