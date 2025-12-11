func largestInteger(nums []int, k int) int {
    n := len(nums)
    if k == n {
        return slices.Max(nums)
    } 
    cnt := make(map[int]int)
    for _, x := range nums {
        cnt[x]++
    }

    if k == 1 {
        mx := -1
        for k, v := range cnt {
            if v == 1 {
                mx = max(mx, k)
            }
        }

        return mx
    }
        
    a, b := max(nums[0], nums[n - 1]), min(nums[0], nums[n - 1])
    if cnt[a] == 1 {
        return a
    } else if cnt[b] == 1 {
        return b
    }
    return -1
}