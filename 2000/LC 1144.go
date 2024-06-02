func movesToMakeZigzag(nums []int) int {    
    f := [2]int{}
    for i, x := range nums {
        l, r := math.MaxInt32, math.MaxInt32
        if i > 0 {
            l = nums[i-1]
        } 
        if i < len(nums) - 1 {
            r = nums[i+1]
        }
        f[i & 1] += max(0,  x - min(l, r) + 1) 
    }

    return min(f[0], f[1])
}
