func garbageCollection(garbage []string, travel []int) int {
    ans := len(garbage[0])
    n := len(garbage)
    m := map[rune]struct{}{}
    for i := n - 1; i > 0; i-- {
        for _, c := range garbage[i] {
            m[c] = struct{}{}
        }
        ans += len(garbage[i]) + len(m) * travel[i-1]
    }

    return ans
}