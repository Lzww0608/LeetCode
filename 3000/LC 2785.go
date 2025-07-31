
import "sort"
func sortVowels(s string) string {
    m := map[byte]bool {
        'a': true, 
        'e': true,
        'i': true,
        'o': true,
        'u': true,
        'A': true, 
        'E': true,
        'I': true,
        'O': true,
        'U': true,
    }

    a := []byte{}
    for i := range s {
        if m[s[i]] {
            a = append(a, s[i])
        }
    }
    sort.Slice(a, func(i, j int) bool {
        return a[i] < a[j]
    })

    ans := make([]byte, 0, len(s))
    j := 0
    for i := range s {
        if m[s[i]] {
            ans = append(ans, a[j])
            j++
        } else {
            ans = append(ans, s[i])
        }
    }

    return string(ans)
}