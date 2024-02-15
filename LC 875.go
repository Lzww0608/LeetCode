func minEatingSpeed(piles []int, h int) int {
    l, r := 1, int(1e9)
    for l < r {
        mid := l + ((r - l) >> 1)
        t := 0
        for _, x := range piles {
            t += (x + mid - 1) / mid
        }
        if t > h {
            l = mid + 1
        } else {
            r = mid
        }
    }
    return l
}