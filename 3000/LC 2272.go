const N int = 26
func largestVariance(s string) (ans int) {

    for i := 0; i < N; i++ {
        for j := 0; j < N; j++ {
            cur := 0
            exist, cnt := false, 0
            for k := range s {
                x := int(s[k] - 'a')
                if x == i {
                    cur++
                } else if x == j {
                    cur--
                    exist = true
                    cnt++
                }

                if cur < 0 {
                    cur, exist = 0, false
                } 

                if exist {
                    ans = max(ans, cur)
                } else if cnt > 0 {
                    ans = max(ans, cur - 1)
                }
            }
        }
    }

    return ans
}