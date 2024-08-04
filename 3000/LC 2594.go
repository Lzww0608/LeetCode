func repairCars(ranks []int, cars int) int64 {
    l, r := 1, 100 * cars * cars


    for l < r {
        mid := l + ((r - l) >> 1)
        sum := 0
        for _, x := range ranks {
            t := float64(mid) / float64(x)
            sum += int(math.Sqrt(t))
        }
        if sum >= cars {
            r = mid
        } else {
            l = mid + 1
        }
    }

    return int64(l)
}