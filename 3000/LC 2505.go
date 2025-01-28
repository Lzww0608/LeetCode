const N int = 60
func subsequenceSumOr(nums []int) int64 {
    cnt := make([]int, N)
    for _, x := range nums {
        for j := 0; (1 << j) <= x; j++ {
            cnt[j] += (x >> j) & 1
        }
    }

    ans := 0
    for i := 0; i < N; i++ {
        if cnt[i] > 0 {
            ans |= 1 << i 
            if i + 1 < N {
                cnt[i+1] += cnt[i] / 2
            }
        }
    }

    return int64(ans)
}