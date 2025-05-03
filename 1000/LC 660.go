func newInteger(n int) int {
    s := strconv.FormatInt(int64(n), 9)
    ans, _ := strconv.Atoi(s)
    return ans
}