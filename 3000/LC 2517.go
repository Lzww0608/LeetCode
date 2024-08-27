func maximumTastiness(price []int, k int) (ans int) {
    sort.Ints(price)
    n := len(price)
    l, r := 0, (price[n-1]- price[0]) / (k - 1) + 1

    check := func(mid int) bool {
        cnt, cur := 1, price[0]
        for _, x := range price  {
            if x >= cur + mid {
                cur = x
                cnt++
            }
        }

        return cnt >= k
    }

    for l < r {
        mid := l + ((r - l) >> 1)
        if check(mid) {
            ans = mid
            l = mid + 1
        } else {
            r = mid
        }
    }

    return 
}