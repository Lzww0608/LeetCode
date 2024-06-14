func minEatingSpeed(piles []int, h int) int {
    l, r := 1, int(1e9)
    for l < r {
        mid := l + ((r - l) >> 1)
        cnt := 0
        for _, x := range piles {
            cnt += (x + mid - 1) / mid
            if cnt > h {
                break
            }
        }
        if cnt > h {
            l = mid + 1
        } else {
            r = mid
        }
    }

    return l
}