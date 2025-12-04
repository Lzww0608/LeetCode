func removeSubstring(s string, k int) string {
    n := len(s)
    byteSt := []byte{}
    cntSt := []int{}

    for i := range n {
        byteSt = append(byteSt, s[i])

        cnt := 1
        if len(byteSt) > 1 && byteSt[len(byteSt) - 2] == s[i] {
            cnt = cntSt[len(cntSt) - 1] + 1
        }
        cntSt = append(cntSt, cnt)

        if s[i] == ')' && cnt == k {
            l := len(byteSt) - k - 1
            if l >= 0 && byteSt[l] == '(' && cntSt[l] >= k {
                byteSt = byteSt[:len(byteSt) - k * 2]
                cntSt = cntSt[:len(cntSt) - k * 2]
            }
        }
    }

    return string(byteSt)
}