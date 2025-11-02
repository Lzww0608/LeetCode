func captureForts(forts []int) (ans int) {
    n := len(forts)

    pre := n
    for i := range n {
        if forts[i] == 1 {
            pre = i
        } else if forts[i] == -1 {
            ans = max(ans, i - pre - 1)
            pre = n
        }
    }

    pre = -1
    for i := n - 1; i >= 0; i-- {
        if forts[i] == 1 {
            pre = i
        } else if forts[i] == -1 {
            ans = max(ans, pre - i - 1)
            pre = -1
        }
    }

    return
}