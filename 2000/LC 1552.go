func maxDistance(position []int, m int) int {
    sort.Ints(position)

    l, r := 0, int(1e9)

    for l + 1 < r {
        mid := l + ((r - l) >> 1)
        res := 1

        for i, j := 1, 0; i < len(position); i++ {
            if mid <= position[i] - position[j] {
                res++
                j = i
            }
        }
        
        if res >= m {
            l = mid
        } else {
            r = mid
        }
    }

    return l
}
