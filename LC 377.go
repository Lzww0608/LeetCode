func combinationSum4(nums []int, target int) int {
    sort.Ints(nums)

    f := make([]int, target + 1)
    f[0] = 1

    for i := 1; i <= target; i++ {
        for _, x := range nums {
            if i >= x {
                f[i] += f[i-x]
            } else {
                break
            }
        }
    }

    return f[target]
}