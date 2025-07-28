func targetIndices(nums []int, target int) (ans []int) {
    sort.Ints(nums)
    for i, x := range nums {
        if x == target {
            ans = append(ans, i)
        } else if x > target {
            break
        }
    }

    return 
}