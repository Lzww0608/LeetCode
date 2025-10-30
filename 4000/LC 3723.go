func maxSumOfSquares(num int, sum int) string {
    if sum > num * 9 {
        return ""
    }

    ans := make([]byte, num)
    for i := range num {
        if sum >= 9 {
            ans[i] = '9'
            sum -= 9
        } else {
            ans[i] = byte('0' + sum)
            sum = 0
        }
    }

    return string(ans)
}