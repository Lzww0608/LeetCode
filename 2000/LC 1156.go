func maxRepOpt1(text string) (ans int) {
    m := map[byte]int{}
    for i := range text {
        m[text[i]]++
    }

    n := len(text)
    l, r := 0, 0
    for l < n {
        r = l 
        for r < n && text[r] == text[l] {
            r++
        }
        k := r + 1
        for k < n && text[k] == text[l] {
            k++
        }

        ans = max(ans, min(m[text[l]], k - l))

        l = r
    }

    return
}
