func distributeCandies(candyType []int) int {
    n := len(candyType) / 2
    m := map[int]bool{}
    for _, x := range candyType {
        m[x] = true
    }

    return min(n, len(m))

}
