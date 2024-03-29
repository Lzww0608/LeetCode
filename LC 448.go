func findDisappearedNumbers(nums []int) (ans []int) {
    n := len(nums)
    for _, x := range nums {
        t := (x - 1) % n 
        nums[t] += n
    }

    for i, x := range nums {
        if x <= n {
            ans = append(ans, i + 1)
        }
    }

    return 
}