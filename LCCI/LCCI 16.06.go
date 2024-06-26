func smallestDifference(a []int, b []int) int {
    sort.Ints(a)
    sort.Ints(b)

    ans := math.MaxInt
    i, j := len(a) - 1, len(b) - 1
    for i >= 0 && j >= 0 {
        if a[i] >= b[j] {
            ans = min(ans, a[i] - b[j])
            i--
        } else {
            ans = min(ans, b[j] - a[i])
            j--
        }
    }

    return ans
}