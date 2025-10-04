const N int = 26
func smallestPalindrome(s string) string {
    cnt := [N]int{}
    for _, c := range s {
        x := int(c - 'a')
        cnt[x]++
    }

    ans := make([]byte, len(s))
    k := 0
    for i := range N {
        for j := 0; j + 1 < cnt[i]; j += 2 {
            ans[k] = byte('a' + i)
            k++
        }

        if cnt[i] & 1 == 1 {
            ans[len(s) / 2] = byte('a' + i)
        }
    }

    for i := len(ans) - 1; i > len(ans) - i - 1; i-- {
        ans[i] = ans[len(ans) - i - 1]
    }

    return string(ans)
}