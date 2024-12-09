func repeatedStringMatch(a string, b string) int {
    m, n := len(a), len(b)
    pi := make([]int, n)
    cnt := 0

    for i := 1; i < n; i++ {
        for cnt > 0 && b[i] != b[cnt] {
            cnt = pi[cnt-1]
        }
        if b[i] == b[cnt] {
            cnt++
        }
        pi[i] = cnt
    }

    cnt = 0
    for i := 0; i < max(m, n) * 2; i++ {
        for cnt > 0 && a[i%m] != b[cnt] {
            cnt = pi[cnt-1]
        }
        if a[i%m] == b[cnt] {
            cnt++
        }
        if cnt == n {
            return (i + m) / m
        }
    }

    return -1
}