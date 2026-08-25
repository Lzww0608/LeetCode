const N int = 100_001

var F [N][]int

func init() {
    for i := 2; i < N; i++ {
        if len(F[i]) != 0 {
            continue
        }
        for j := i; j < N; j += i {
            F[j] = append(F[j], i)
        }
    }
}

func longestSubarray(nums []int, k int) (ans int) {
    n := len(nums)
    cnt := make(map[int]int)
    for l, r := 0, 0; r < n; r++ {
        for _, y := range F[nums[r]] {
            cnt[y]++
        }

        for len(cnt) > k {
            for _, y := range F[nums[l]] {
                if cnt[y]--; cnt[y] == 0 {
                    delete(cnt, y)
                }
            }
            l++
        }

        ans = max(ans, r - l + 1)
    }

    return
}