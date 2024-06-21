func maxAlternatingSum(nums []int) int64 {
    var odd, even int64

    for _, x := range nums {
        odd, even = max(odd, even - int64(x)), max(even, odd + int64(x))
    }

    return max(odd, even)
}
