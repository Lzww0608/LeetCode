func getPermutation(n int, k int) string {
    s := make([]byte, n)
    for i := range s {
        s[i] = byte('1' + i)
    }

    for i := 1; i < k; i++ {
        nextPermutation(s)
    }

    return string(s)
}

func nextPermutation(s []byte) {
    n := len(s)
    i := n - 2
    for i >= 0 && s[i] > s[i+1] {
        i--
    }
    if i < 0 {
        return
    }
    j := n - 1
    for j > 0 && s[j] < s[i] {
        j--
    }
    s[i], s[j] = s[j], s[i]
    for p, q := i + 1, n - 1; p < q; p, q = p + 1, q - 1 {
        s[p], s[q] = s[q], s[p]
    }
}