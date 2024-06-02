func carPooling(trips [][]int, capacity int) bool {
    m := make([]int, 1005)
    for _, t := range trips {
        m[t[1]] += t[0]
        m[t[2]] -= t[0]
    }
    sum := 0
    for _, x := range m {
        sum += x
        if sum > capacity {
            return false
        }
    }
    return sum <= capacity
}
