func smallestDivisor(nums []int, threshold int) int {
    l, r := 1, int(1e6)

    cal := func(div int) int {
        res := 0
        for _, x := range nums {
            res += (x + div - 1) / div
        }
        return res
    }

    for l <= r {
        mid := l + ((r - l) >> 1)
        if cal(mid) > threshold {
            l = mid + 1
        } else {
            r = mid - 1
        }
    }

    return l
}
