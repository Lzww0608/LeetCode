func recoverArray(nums []int) []int {
    cnt := make(map[int]int)
    for _, x := range nums {
        cnt[x]++
    }
    sort.Ints(nums)
    n := len(nums)

    solve := func(k int) ([]int, bool) {
        a := make([]int, 0, n / 2)
        cur := make(map[int]int)
        for _, x := range nums {
            if cur[x] == cnt[x] {
                continue
            }
            cur[x]++
            y := x + k * 2 
            if cur[y]++; cur[y] > cnt[y] {
                return nil, false
            }
            a = append(a, x + k)
        }
        return a, true
    }

    for i := 1; i < n; i++ {
        if (nums[i] - nums[0]) & 1 == 1 || nums[i] - nums[0] == 0 {
            continue
        }
        k := (nums[i] - nums[0]) / 2
        if a, ok := solve(k); ok {
            return a
        }
    }

    return nil
}