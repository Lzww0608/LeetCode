func minMoves2(nums []int) (ans int) {
    n := len(nums)
    sort.Ints(nums)

    y := nums[n / 2]
    for _, x := range nums {
        ans += abs(x - y)
    }

    return 
}

func abs(x int) int {
    if x < 0 {
        return -x
    }

    return x
}