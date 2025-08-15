func checkArithmeticSubarrays(nums []int, l []int, r []int) []bool {
    m := len(l)
    ans := make([]bool, m)
    
    for i := 0; i < m; i++ {
        left, right := l[i], r[i]
        a := make([]int, right - left + 1)
        copy(a, nums[left : right + 1])
        sort.Ints(a)

        f := true 
        for i := 2; i < len(a); i++ {
            if a[i] - a[i - 1] != a[i - 1] - a[i - 2] {
                f = false
                break
            }
        }
        ans[i] = f
    }

    return ans
}