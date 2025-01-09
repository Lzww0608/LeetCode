func evenProduct(nums []int) int64 {
    ans := 0
    l := 0

    sum := 0
    for _, x := range nums {
        if x & 1 == 0 {
            sum++
        }

        for sum > 0 {
            if nums[l] & 1 == 0 {
                sum--
            }
            l++
        }

        ans += l
    }

    return int64(ans)
}