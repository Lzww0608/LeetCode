func smallestDistancePair(nums []int, k int) int {
    l, r := 0, int(1e6)
    sort.Ints(nums)
    n := len(nums)

    cal := func(target int) (res int) {
        i, j := 0, 1
        for i < n {
            for j < n && nums[j] - nums[i] <= target {
                j++
            }
            res += j - i - 1
            i++
        }
        return res
    }

    for l < r {
        mid := l + ((r - l) >> 1)
        cnt := cal(mid)
        if cnt < k {
            l = mid + 1
        } else {
            r = mid 
        }
    }

    return l
}