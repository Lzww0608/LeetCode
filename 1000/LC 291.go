
import "strings"
func wordPatternMatch(pattern string, s string) bool {
    m, n := len(pattern), len(s)
    if m > n {
        return false
    }
    mp := map[byte]string{}
    vis := map[string]bool{}

    var dfs func(int, int) bool
    dfs = func(i, j int) bool {
        if i >= m && j >= n {
            return true 
        } else if i >= m || j >= n {
            return false
        }

        c := pattern[i]
        if v, ok := mp[c]; ok {
            pos := strings.Index(s[j:], v) + j
            if pos != j {
                return false
            }
            return dfs(i + 1, j + len(v))
        } else {
            for k := 1; j + k <= n - (m - i - 1); k++ {
                
                if vis[s[j:j+k]] {
                    continue
                }
                mp[c] = s[j:j+k]
                vis[s[j:j+k]] = true
                if dfs(i + 1, j + k) {
                    return true
                }
                delete(mp, c)
                delete(vis, s[j:j+k])
            }
        }

        return false
    }

    return dfs(0, 0)
}