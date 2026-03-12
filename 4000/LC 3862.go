func smallestBalancedIndex(nums []int) int {
    sum := 0

    for _, x := range nums {
        sum += x
    }

    n := len(nums)
    prod := 1
    for i := n - 1; i >= 0; i-- {
        sum -= nums[i]
        if sum == prod {
            return i
        } else if prod > sum / nums[i] {
            break
        }
        prod *= nums[i]
    }

    return -1
}