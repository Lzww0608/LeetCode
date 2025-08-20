func maxDepthAfterSplit(seq string) []int {
    n := len(seq)
    ans := make([]int, n)
    a, b := 0, 0
    for i := range seq {
        if seq[i] == '(' {
            if a > b {
                ans[i] = 1
                b++
            } else {
                a++
            }
        } else {func largestPerimeter(nums []int) int64 {
    sort.Ints(nums)
    n := len(nums)
    sum := 0
    for _, x := range nums {
        sum += x
    }

    for i := n - 1; i >= 0; i-- {
        x := nums[i]
        if sum - x > x {
            return int64(sum)
        }
        sum -= x
    }

    return -1
}
            if a > b {
                a--
            } else {
                ans[i] = 1
                b--
            }
        }
    }

    return ans
}