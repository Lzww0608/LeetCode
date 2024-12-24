func minMovesToMakePalindrome(str string) (ans int) {
    s := []byte(str)
    n := len(s)
next:
    for i, j := 0, n - 1; i < j; i++ {
        for k := j; k > i; k-- {
            if s[i] == s[k] {
                for ; k < j; k++ {
                    s[k], s[k+1] = s[k+1], s[k]
                    ans++
                }
                j--
                continue next
            }
        }
        ans += n / 2 - i
    }

    return
}