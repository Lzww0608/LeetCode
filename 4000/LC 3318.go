func findXSum(nums []int, k int, x int) []int {
    n := len(nums)
    ans := make([]int, n - k + 1)

    a := []int{}
    for i := range n - k + 1 {
        cnt := make(map[int]int)
        for j := 0; j < k; j++ {
            cnt[nums[i + j]]++
        }
        a := a[:0]
        for k := range cnt {
            a = append(a, k)
        }
        sort.Slice(a, func(i, j int) bool {
            return cnt[a[i]] > cnt[a[j]] || cnt[a[i]] == cnt[a[j]] && a[i] > a[j]
        })

        for j := 0; j < x && j < len(a); j++ {
            ans[i] += a[j] * cnt[a[j]]
        }
    }

    return ans
}