
import "math/bits"
func tripletCount(a []int, b []int, c []int) int64 {
    cnt := make([]int, 2)
    for _, x := range a {
        if bits.OnesCount(uint(x)) & 1 == 0 {
            cnt[0]++
        } else {
            cnt[1]++
        }
    }

    cnt1 := make([]int, 2)
    for _, x := range b {
        if bits.OnesCount(uint(x)) & 1 == 0 {
            cnt1[0] += cnt[0]
            cnt1[1] += cnt[1]
        } else {
            cnt1[0] += cnt[1]
            cnt1[1] += cnt[0]
        }
    }

    ans := 0
    for _, x := range c {
        if bits.OnesCount(uint(x)) & 1 == 0 {
            ans += cnt1[0]
        } else {
            ans += cnt1[1]
        }
    }

    return int64(ans)
}