func magicalString(n int) (cnt int) {
    ans := make([]int, n + 1)
    ans[0], ans[1] = 1, 2
    for i, j := 1, 1; i < n && j < n; i++ {
        if ans[i] == 2 {
            ans[j] = ans[j - 1] ^ 3
            ans[j + 1] = ans[j]
            j += 2
        } else {
            ans[j] = ans[j - 1] ^ 3
            j++
        }
    }

    for i := 0; i < n; i++ {
        if ans[i] == 1 {
            cnt++
        }
    }

    return cnt
}