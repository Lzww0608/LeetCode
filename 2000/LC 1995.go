func countQuadruplets(nums []int) (ans int) {
    n := len(nums)
    cnt := make(map[int]int)

    for b := n - 3; b >= 1; b-- {
        for d := b + 2; d < n; d++ {
            cnt[nums[d] - nums[b+1]]++
        } 

        for a := b - 1; a >= 0; a-- {
            ans +=  cnt[nums[a] + nums[b]]
        }
    }

    return 
}