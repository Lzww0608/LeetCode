func kthLargestNumber(nums []string, k int) string {
    compare := func(a, b string) bool {
        if len(a) != len(b) {
            return len(a) > len(b)
        }

        for i := range a {
            if a[i] != b[i] {
                return a[i] > b[i]
            }
        }

        return true
    }

    sort.Slice(nums, func(i, j int) bool {
        return compare(nums[i], nums[j])
    })

    return nums[k-1]
}