func digitSum(s string, k int) string {
    ans := []byte(s)
    for len(ans) > k {
        tmp := []byte{}
        cur := 0
        for i := range ans {
            cur += int(ans[i] - '0')
            if i % k == k - 1 || i == len(ans) - 1 {
                tmp = append(tmp, strconv.Itoa(cur)...)
                cur = 0
            }
        }
        ans = tmp
    }

    return string(ans)
}