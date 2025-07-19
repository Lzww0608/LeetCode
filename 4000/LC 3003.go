
import (
	"math"
	"math/bits"
)
func maxPartitionsAfterOperations(s string, k int) int {
    memo := make(map[int]int)

    var dfs func(int, int, bool) int 
    dfs = func(i, mask int, changed bool) (res int) {
        if i == len(s) {
            return 1
        }

        t := i << 30 | mask << 1 
        if changed {
            t |= 1
        }
        if v, ok := memo[t]; ok {
            return v
        }
        
        res = math.MinInt32
        x := 1 << int(s[i] - 'a')
        if !changed {
            for j := 0; j < 26; j++ {
                cur := mask | (1 << j) 
                if bits.OnesCount(uint(cur)) > k {
                    res = max(res, dfs(i + 1, 1 << j, true) + 1)
                } else {
                    res = max(res, dfs(i + 1, cur, true))
                }
            }
            
        }
        
        if bits.OnesCount(uint(x | mask)) > k {
            res = max(res, dfs(i + 1, x, changed) + 1)
        } else {
            res = max(res, dfs(i + 1, mask | x, changed))
        }

        memo[t] = res  
        return  
    }

    return dfs(0, 0, false)
}