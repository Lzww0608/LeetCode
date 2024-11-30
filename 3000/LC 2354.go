import "math/bits"
func countExcellentPairs(nums []int, k int) int64 {
    ans := 0
    cnt := make([]int, 64)
    m := map[int]bool{}
    for _, x := range nums {
        y := bits.OnesCount(uint(x))
        if !m[x] {
            cnt[y]++
            m[x] = true
        } 
    }
    pre := make([]int, len(cnt))
    for i := len(cnt) - 2; i >= 0; i-- {
        pre[i] = pre[i+1] + cnt[i]
    }

    for i := range cnt {
        if cnt[i] == 0 {
             continue
        }
        ans += pre[max(0, k - i)] * cnt[i]
        //fmt.Println(i, ans)
    }

    return int64(ans)
}