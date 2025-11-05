func alternateDigitSum(n int) int {
    f := 1
    ans := 0

    s := strconv.Itoa(n)
    for i := range s {
        x := int(s[i] - '0')
        if f == 1 {
            ans += x
        } else {
            ans -= x
        }
        f ^= 1
    }

    return ans
}