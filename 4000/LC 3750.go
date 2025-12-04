func minimumFlips(n int) (ans int) {
    s := strconv.FormatInt(int64(n), 2)
    m := len(s)

    for i := range m {
        if s[i] != s[m - i - 1] {
            ans++
        }
    }

    return 
}