func sumOfDigits(nums []int) int {
    x := slices.Min(nums)
    sum := 0
    for x > 0 {
        sum += x % 10
        x /= 10
    }

    return (sum & 1) ^ 1
}