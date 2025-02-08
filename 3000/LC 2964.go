func divisibleTripletCount(nums []int, d int) (ans int) {
    cnt := make(map[int]int)
    n := len(nums)
    for i := 1; i < n; i++ {
        ans += cnt[(d-nums[i]%d)%d]
        for j := 0; j < i; j++ {
            cnt[(nums[i] + nums[j])%d]++
        }
    }

    return
}