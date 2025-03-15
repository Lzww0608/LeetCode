
import "slices"
func movesToStamp(stamp string, target string) (ans []int) {
    n, m := len(target), len(stamp)
    s := []byte(stamp)
    t := []byte(target)

    check := func(idx int) {
        cnt := 0
        for i, j := idx, 0; i < n && j < m; i, j = i + 1, j + 1 {
            if s[j] != t[i] && t[i] != '#' {
                return
            }
        } 

        for i, j := idx, 0; i < n && j < m; i, j = i + 1, j + 1 {
            if t[i] != '#' {
                cnt++
            }
            t[i] = '#'
        } 
        if cnt > 0 {
            ans = append(ans, idx)
        }
        
    }

    for i := 0; i <= n - m; i++ {
        check(i)
    }

    for i := n - m; i >= 0; i-- {
        check(i)
    }

    for i := range t {
        if t[i] != '#' {
            return []int{}
        }
    }

    slices.Reverse(ans)

    return
}