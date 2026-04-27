func kthRemainingInteger(nums []int, queries [][]int) []int {
    pos := []int{}
    for i, x := range nums {
        if x & 1 == 0 {
            pos = append(pos, i)
        }
    }

    ans := make([]int, len(queries))
    for i, q := range queries {
        l0 := sort.SearchInts(pos, q[0])
        r0 := sort.SearchInts(pos, q[1] + 1)
        k := q[2]

        l, r := -1, r0 - l0
        for l + 1 < r {
            mid := l + ((r - l) >> 1)
            if nums[pos[l0 + mid]] / 2 - 1 - mid >= k {
                r = mid
            } else {
                l = mid
            }
        }

        ans[i] = (r + k) * 2
    }

    return ans
}