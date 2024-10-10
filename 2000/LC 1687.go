func boxDelivering(boxes [][]int, portsCount int, maxBoxes int, maxWeight int) int {
    n := len(boxes)
    diff := make([]int, n + 1)
    w := make([]int, n + 1)
    for i := 1; i <= n; i++ {
        w[i] = w[i-1] + boxes[i-1][1]
        if i > 1 {
            if boxes[i-2][0] != boxes[i-1][0] {
                diff[i] = diff[i-1] + 1
            } else {
                diff[i] = diff[i-1]
            }
        }
    }

    st := []int{0}
    f := make([]int, n + 1)
    g := make([]int, n + 1)
    for i := 1; i <= n; i++ {
        for len(st) > 0 && ((i - st[0]) > maxBoxes || w[i] - w[st[0]] > maxWeight) {
            st = st[1:]
        }

        f[i] = g[st[0]] + diff[i] + 2
        
        if i != n {
            g[i] = f[i] - diff[i+1]
            for len(st) > 0 && g[i] <= g[st[len(st)-1]] {
                st = st[:len(st)-1]
            }
            st = append(st, i)
        }
    }

    return f[n]
}