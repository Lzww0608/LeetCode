
import "unicode"
func countValidWords(sentence string) (ans int) {
    str := strings.Split(sentence, " ")
next:
    for _, s := range str {
        if s == "" {
            continue
        }
        n := len(s)
        cnt := 0 
        for i, c := range s {
            if unicode.IsDigit(c) {
                continue next
            }
            if c == '-' {
                if i == 0 || i == n - 1 || !unicode.IsLower(rune(s[i-1])) || !unicode.IsLower(rune(s[i+1])) {
                        continue next
                    }
                cnt++
            } else if !unicode.IsLower(c) {
                if i != n - 1 {
                    continue next
                }
            }
        }
        if cnt <= 1 {
            ans++
        }
    }

    return
}