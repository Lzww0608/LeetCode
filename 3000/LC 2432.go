func hardestWorker(n int, logs [][]int) int {
    ans, mx := 0, 0

    pre := 0
    for _, log := range logs {
        i, t := log[0], log[1]
        x := t - pre
        if x > mx || x == mx && i < ans {
            ans, mx = i, x
        }
        pre = t
    }

    return ans
}