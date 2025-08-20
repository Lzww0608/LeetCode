func minimumPerimeter(neededApples int64) int64 {
    l, r := 1, 100_000
    for l < r {
        mid := l + ((r - l) >> 1)
        sum := mid * 2 * (mid + 1) * (2 * mid + 1)
        if int64(sum) < neededApples {
            l = mid + 1
        } else {
            r = mid
        }
    }

    return int64(l) * 8
}