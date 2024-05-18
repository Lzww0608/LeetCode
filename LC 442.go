func findDuplicates(nums []int) (ans []int) {
    for _, x := range nums {
        x = abs(x)
        if nums[x-1] > 0 {
            nums[x-1] = -nums[x-1]
        } else {
            ans = append(ans, x)
        }
    }

    return
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}