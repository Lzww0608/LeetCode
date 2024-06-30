func largestPerimeter(nums []int) (ans int) {
    sort.Ints(nums)
    n := len(nums)
    
    for i := n - 3; i >= 0; i-- {
        a, b, c := nums[i], nums[i+1], nums[i+2]
        if a + b > c {
            ans = a + b + c
            break
        }
    }

    return
}
