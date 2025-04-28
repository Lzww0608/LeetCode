func minimumOperations(nums []int, target []int) int64 {
    n := len(nums)
    a := make([]int, n)
    a[0] = nums[0] - target[0]
    for i := 1; i < n; i++ {
        a[i] = nums[i] - nums[i-1] - (target[i] - target[i-1])
    }

    pos, neg := 0, 0 
    for _, x := range a {
        if x < 0 {
            neg -= x
        } else {
            pos += x 
        }
    }

    return int64(max(pos, neg))
}