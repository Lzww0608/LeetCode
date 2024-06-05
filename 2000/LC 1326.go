func minTaps(n int, ranges []int) int {
    rightMost := make([]int, n + 1)
    for i, x := range ranges {
        left := max(0, i - x)
        rightMost[left] = max(rightMost[left], i + x)
    }

    cur, r := 0, 0
    ans := 0
    for i, x := range rightMost[:n] {
        r = max(r, x)
        if i == cur {
            if cur == r {
                return -1
            }
            cur = r
            ans++
        }
    }

    return ans
}
