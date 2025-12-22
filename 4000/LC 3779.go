func minOperations(nums []int) (ans int) {
    cnt := make(map[int]int)
    sum := 0
    for _, x := range nums {
        if cnt[x]++; cnt[x] == 2 {
            sum++
        }
    }

    for sum > 0 && len(nums) > 0 {
        ans++
        for i := 0; len(nums) > 0 && i < 3; i++ {
            x := nums[0]
            nums = nums[1:]
            if cnt[x]--; cnt[x] == 1 {
                sum--
            }
        }
    }

    return
}