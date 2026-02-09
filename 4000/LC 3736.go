func minMoves(nums []int) (ans int) {
    mx := slices.Max(nums)
    for _, x := range nums {
        ans += mx - x
    }

    return 
}