func longestBeautifulSubstring(word string) (ans int) {
    l, r, n := 0, 1, len(word)
    m := map[byte]int{
        'a': 0,
        'e': 1,
        'i': 2,
        'o': 3,
        'u': 4,
    }
    cnt := [5]int{}

    check := func() bool {
        for _, x := range cnt {
            if x == 0 {
                return false
            }
        }
        return true
    }

    cnt[m[word[0]]]++
    for r < n {
        if word[r] >= word[r-1] {
            cnt[m[word[r]]]++
            if check() {
                ans = max(ans, r - l + 1)
            }
        } else {
            l = r
            for i := range cnt {
                cnt[i] = 0
            }
            cnt[m[word[r]]]++
        }
        r++
    }

    return
}
