func lengthOfLIS(nums []int) (ans int) {
    st := []int{math.MinInt32}

    for _, x := range nums {
        if x <= st[len(st)-1] {
            pos := sort.SearchInts(st, x)
            st[pos] = x
        } else {
            st = append(st, x)
        }

        ans = max(ans, len(st) - 1)
    }

    return ans
}