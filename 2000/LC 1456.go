func maxVowels(s string, k int) (ans int) {
    vowel := 0b100000100000100010001

    cnt := 0
    for i := range s {
        if vowel & (1 << int(s[i] - 'a')) != 0 {
            cnt++
        }

        ans = max(ans, cnt)
        if i >= k - 1 {
            if vowel & (1 << int(s[i - k + 1] - 'a')) != 0 {
                cnt--
            }
        }
    }

    return
}