func leastOpsExpressTarget(x int, target int) int {
    a := []int{}
    for target > 0 {
        a = append(a, target % x)
        target /= x 
    }

    n := len(a)
    add := make([]int, n)
    sub := make([]int, n)
    add[0], sub[0] = a[0] * 2, (x - a[0]) * 2
    for i := 1; i < n; i++ {
        add[i] = min(i * a[i] + add[i-1], i * (a[i] + 1) + sub[i-1])
        sub[i] = min(i * (x - a[i]) + add[i-1], i * (x - a[i] - 1) + sub[i-1])
    }

    return min(add[n-1], sub[n-1] + n) - 1
}