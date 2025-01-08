
import "strings"
const N int = 26
func longestNiceSubstring(s string) (ans string) {
    t := strings.ToLower(s)
    n := len(s)
    for k := 1; k <= N; k++ {
        upper := make([]int, N)
        lower := make([]int, N)
        cnt := make(map[int]int)
        for l, r := 0, 0; r < n; r++ {
            c := s[r]
            x := int(t[r] - 'a')
            if c >= 'a' && c <= 'z' {
                lower[x]++
                cnt[x] |= 1
            } else {
                upper[x]++
                cnt[x] |= 2
            }

            for len(cnt) > k {
                c = s[l]
                x := int(t[l] - 'a')
                if c >= 'a' && c <= 'z' {
                    if lower[x]--; lower[x] == 0 {
                        cnt[x] &^= 1
                    }
                } else {
                    if upper[x]--; upper[x] == 0 {
                        cnt[x] &^= 2
                    }
                }
                if cnt[x] == 0 {
                    delete(cnt, x)
                }
                l++
            }

            sum := 0
            for _, x := range cnt {
                if x == 3 {
                    sum++
                }
            }
            if sum == len(cnt) && r - l + 1 > len(ans) {
                ans = s[l:r+1]
            }
        }
    } 
    
    return
}