func validPalindrome(s string) bool {
    n := len(s)
    cnt := 0

    var check func(int, int) bool
    check = func(l, r int) bool {
        for ; l < r; l, r = l + 1, r - 1 {
            if s[l] != s[r] {
                cnt++
                if cnt > 1 {
                    return false
                }
                return check(l + 1, r) || check(l, r - 1)
            }
        }

        return true
    }

    return check(0, n - 1)
}