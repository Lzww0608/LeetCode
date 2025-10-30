func sortPermutation(nums []int) int {
    if slices.IsSorted(nums) {
        return 0
    }
    
    n := len(nums)
    a := make([]int, n)
    copy(a, nums)
    sort.Ints(a)

    ans := math.MaxInt32 
    for i, x := range nums {
        if x != a[i] {
            ans &= x
        }
    }

    return ans
}