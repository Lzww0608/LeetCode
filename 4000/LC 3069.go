func resultArray(nums []int) []int {
    a := []int{nums[0]}
    b := []int{nums[1]}
    for i := 2; i < len(nums); i++ {
        if a[len(a) - 1] >= b[len(b) - 1] {
            a = append(a, nums[i])
        } else {
            b = append(b, nums[i])
        }
    }

    a = append(a, b...)
    return a
}