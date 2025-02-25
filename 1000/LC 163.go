func findMissingRanges(nums []int, lower int, upper int) (ans [][]int) {
    p := lower
    nums = append(nums, upper + 1)
    for _, x := range nums {
        if x == p {
            p++
        } else {
            ans = append(ans, []int{p, x - 1})
            p = x + 1 
        }
    }

    return
}