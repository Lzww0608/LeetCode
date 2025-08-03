func smallestNumber(n int64) string {
    if n == 1 {
        return "1"
    }
    ans := []byte{}
    for i := int64(9); i > 1; i-- {
        for n % i == 0 {
            ans = append(ans, byte('0' + i))
            n /= i
        }
    }

    if n > 1 {
        return "-1"
    }

    sort.Slice(ans, func(i, j int) bool {
        return ans[i] < ans[j]
    })

    return string(ans)
}