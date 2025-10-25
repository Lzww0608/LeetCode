(int(ans[i] - '0') + int(ans[i + 1] - '0') % 10)func hasSameDigits(s string) bool {
    ans := []byte(s)

    for len(ans) > 2 {
        tmp := make([]byte, len(ans) - 1)
        for i := 0; i < len(ans) - 1; i++ {
            tmp[i] = byte('0' + (int(ans[i] - '0') + int(ans[i + 1] - '0')) % 10)
            
        }
        ans = tmp
    }

    return ans[0] == ans[1]
}