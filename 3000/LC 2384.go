func largestPalindromic(num string) string {
    cnt := make([]int, 10)
    for _, c := range num {
        x := int(c - '0')
        cnt[x]++
    }

    if cnt[0] == len(num) {
        return "0"
    }

    odd := -1
    for i, x := range cnt {
        if x & 1 == 1 {
            odd = i
        }
    }

    s := []byte{}
    ans := []byte{}
    for i := 9; i > 0 || (i == 0 && len(s) > 0); i-- {
        x := cnt[i]
        for j := 0; j < x / 2; j++ {
            s = append(s, byte('0' + i))
        }
    }

    ans = append(ans, s...)
    if odd != -1 {
        s = append(s, byte('0' + odd))
    }

    var n int = len(s)
    for i := n - 1; i >= 0; i-- {
        ans = append(ans, s[i])
    }


    return string(ans)
}
