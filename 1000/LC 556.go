func nextGreaterElement(n int) int {
    s := []byte(strconv.Itoa(n))
    m := len(s)

    i := m - 2
    for i >= 0 && s[i] >= s[i+1] {
        i--
    }
    if i < 0 {
        return -1
    }

    j := m - 1
    for s[i] >= s[j] {
        j--
    }
    s[i], s[j] = s[j], s[i]

    for p, q := i + 1, m - 1; p < q; p, q = p + 1, q - 1 {
        s[p], s[q] = s[q], s[p]
    }

    ans, _ := strconv.Atoi(string(s))
    if ans > math.MaxInt32 {
        return -1
    }
    return ans
}