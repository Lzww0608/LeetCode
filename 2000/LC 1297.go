func maxFreq(s string, maxLetters int, minSize int, maxSize int) (ans int) {
    m := map[string]int{}
    n := len(s)

    for i := range s {
        j := i
        x := 0
        for j < n && j - i < maxSize {
            x |= 1 << int(s[j] - 'a')
            if bits.OnesCount(uint(x)) > maxLetters {
                break
            }

            if j - i + 1 >= minSize {
                if m[s[i:i+minSize]]++; m[s[i:i+minSize]] > ans {
                    ans = m[s[i:i+minSize]]
                }
                break
            }
            j++
        }
    }

    return
}