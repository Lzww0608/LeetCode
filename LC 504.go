func convertToBase7(num int) string {
    ans, k := 0, 1
    for num != 0 {
        ans += num % 7 * k
        num /= 7
        k *= 10
    }
    return strconv.Itoa(ans)
}