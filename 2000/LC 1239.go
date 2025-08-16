
import "math/bits"
func maxLength(arr []string) (ans int) {
    n := len(arr)
    cnt := make([]int, 0, n)
    for _, s := range arr {
        cur := 0
        for _, c := range s {
            if cur & (1 << int(c - 'a')) != 0 {
                cur = 0
                break
            }
            cur |= 1 << int(c - 'a')
        }
        if cur != 0 {
            cnt = append(cnt, cur)
        }
    }

    n = len(cnt)
    m := 1 << n
    next:
    for i := 1; i < m; i++ {
        cur := 0
        for j := 0; (1 << j) <= i; j++ {
            if i & (1 << j) != 0 {
                if cur & cnt[j] != 0 {
                    continue next
                }
                cur |= cnt[j]
            }
        }
        ans = max(ans, bits.OnesCount(uint(cur)))
    }

    return 
}