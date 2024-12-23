import (
    "strings"
)
func shortestBeautifulSubstring(s string, k int) (ans string) {
    if strings.Count(s, "1") < k {
        return ""
    }

    if k == 1 {
        return "1"
    }

    cnt, l := 0, 0
    pos := []int{}
    for r := range s {
        if s[r] == '1' {
            pos = append(pos, r)
            if cnt++; cnt == k {
                if ans == "" || len(ans) > r - l + 1 || len(ans) == r - l + 1 && ans > s[l:r+1] {
                    ans = s[l:r+1]
                }
                pos = pos[1:]
                if len(pos) > 0 {
                    l = pos[0]
                }
                cnt--
            } else if cnt == 1 {
                l = r
            }
        }
    }

    return
}