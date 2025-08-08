
import "sort"
func sortEvenOdd(nums []int) []int {
    n := len(nums)
    odd := make([]int, 0, n / 2)
    even := make([]int, 0, (n + 1) / 2)
    for i, x := range nums {
        if i & 1 == 0 {
            even = append(even, x)
        } else {
            odd = append(odd, x)
        }
    }

    sort.Ints(even)
    sort.Slice(odd, func(i, j int) bool {
        return odd[i] > odd[j]
    })
    for i := 0; i < n; i++ {
        if i & 1 == 0 {
            nums[i] = even[i / 2]
        } else {
            nums[i] = odd[i / 2]
        }
    }
    return nums
}