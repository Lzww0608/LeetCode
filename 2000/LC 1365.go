const N int = 100
func smallerNumbersThanCurrent(nums []int) []int {
    cnt := make([]int, N + 1)
    for _, x := range nums {
        cnt[x]++
    }

    for i := 1; i <= N; i++ {
        cnt[i] += cnt[i-1]
    }

    for i := range nums {
        if nums[i] != 0 {
            nums[i] = cnt[nums[i]-1]
        }
    }

    return nums
}