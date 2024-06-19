func decodeAtIndex(s string, k int) string {
    j, m := -1, 0
    var ans string
    for i, c := range s {
        if c >= 'a' && c <= 'z' {
            m++
            ans = string(c)
        } else {
            x := int(c - '0')
            m *= x
            if m > k {
                j = i
                break
            }
        }
        if m == k {
            return ans
        }
    }

    for ; j >= 0; j-- {
        k %= m
        if k == 0 && s[j] >= 'a' && s[j] <= 'z' {
            return string(s[j])
        } else if s[j] >= '2' && s[j] <= '9' {
            m /= int(s[j] - '0')
        } else {
            m--
        }
    }

    return ""
}