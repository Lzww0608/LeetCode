func fullJustify(words []string, maxWidth int) (ans []string) {
    r, n := 0, len(words)
    for {
        l, sum := r, 0
        for r < n && sum + len(words[r]) + r - l <= maxWidth {
            sum += len(words[r])
            r++
        }

        if r == n {
            s := strings.Join(words[l:], " ")
            ans = append(ans, s + blank(maxWidth - len(s)))
            return
        }

        cnt := r - l
        space := maxWidth - sum

        if cnt == 1 {
            ans = append(ans, words[l] + blank(space))
            continue
        }
        
        avg := space / (cnt - 1)
        extra := space % (cnt - 1)
        a := strings.Join(words[l:l+extra+1], blank(avg + 1))
        b := strings.Join(words[l+extra+1:r], blank(avg))
        ans = append(ans, a + blank(avg) + b)
    }
}

func blank(n int) string {
    return strings.Repeat(" ", n)
}