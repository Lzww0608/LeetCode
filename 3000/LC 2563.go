func countFairPairs(nums []int, lower int, upper int) (ans int64) {
    sort.Ints(nums)


    find := func(x, t int) int {
        l, r := 0, t
        for l < r {
            mid := l + ((r - l) >> 1)
            if nums[mid] < x {
                l = mid + 1
            } else {
                r = mid
            }
        }
        return l
    }

    for i, x := range nums {
        l, r := find(lower - x, i), find(upper - x + 1, i)
        ans += int64(r) - int64(l)
    }

    return 
}
