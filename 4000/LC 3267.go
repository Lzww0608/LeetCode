
import "strconv"
func countPairs(nums []int) (ans int) {
    sort.Ints(nums)
    cnt := make(map[int]int)
    for _, x := range nums {
        s := []byte(strconv.Itoa(x))
        m := len(s)
        cur := make(map[int]bool)
        cur[x] = true
        for i := 0; i < m; i++ {
            for j := i + 1; j < m; j++ {
                if s[i] == s[j] {
                    continue
                }
                s[i], s[j] = s[j], s[i]
                t, _ := strconv.Atoi(string(s))
                cur[t] = true

                for p := i + 1; p < m; p++ {
                    for q := p + 1; q < m; q++ {
                        s[p], s[q] = s[q], s[p]
                        t, _ = strconv.Atoi(string(s))
                        cur[t] = true
                        s[p], s[q] = s[q], s[p]
                    }
                }

                s[i], s[j] = s[j], s[i]
            }
            
        }
        for k := range cur {
            ans += cnt[k]
        }

        cnt[x]++
    }

    return 
}