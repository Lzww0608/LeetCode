func validPalindrome(s string) bool {
    cnt, n := 0, len(s)
    var check func(int, int) bool
    check = func(l, r int) bool {
        for ; l < r; l, r = l + 1, r - 1 {
            if s[l] != s[r] {
                cnt++
                if cnt > 1 {
                    return false
                }
                if s[l+1] == s[r] {
                    if s[l] == s[r-1] {
                        return check(l+1, r) || check(l, r-1)
                    } 
                    return check(l+1, r)
                } else if s[l] == s[r-1] {
                    return check(l, r-1)
                } else {
                    return false
                }
            }
        }
        return true
    }
    

    return check(0, n - 1)

}
