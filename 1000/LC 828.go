func uniqueLetterString(s string) (ans int) {
    var last0, last1 [26]int
    for i := range last0 {
        last0[i], last1[i] = -1, -1
    }
    cnt := 0
    for i, c := range s {
        x := int(c - 'A')
        cnt += i - last1[x] - (last1[x] - last0[x])
        ans += cnt
        last0[x], last1[x] = last1[x], i
    }

    return
}