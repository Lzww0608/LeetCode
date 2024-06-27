func minIncrements(n int, a []int) (ans int) {
    for i := n / 2; i > 0; i-- {
        mx := max(a[2 * i - 1], a[2 * i])
        ans += abs(a[2 * i - 1] - a[2 * i])
        a[i-1] += mx;
    }
    return
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
