
import "math"
func maxDifference(s string, k int) int {
    ans := math.MinInt32 
    for x := 0; x < 5; x++ {
        for y := 0; y < 5; y++ {
            if x == y {
                continue
            }
            cur_s := [5]int{}
            pre_s := [5]int{}
            min_s := [2][2]int{{math.MaxInt32, math.MaxInt32}, {math.MaxInt32, math.MaxInt32}}
            l := 0
            for i := 0; i < len(s); i++ {
                cur_s[s[i] - '0']++
                r := i + 1
                for r - l >= k && cur_s[x] > pre_s[x] && cur_s[y] > pre_s[y] {
                    p := &min_s[pre_s[x] & 1][pre_s[y] & 1]
                    *p = min(*p, pre_s[x] - pre_s[y])
                    pre_s[s[l] - '0']++
                    l++
                }
                ans = max(ans, cur_s[x] - cur_s[y] - min_s[cur_s[x] & 1 ^ 1][cur_s[y] & 1])
            }
        }
    }

    return ans
}