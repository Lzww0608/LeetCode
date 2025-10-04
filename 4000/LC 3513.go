func uniqueXorTriplets(nums []int) int {
    n := len(nums)
    if n <= 2 {
        return n
    }
    ans := (1 << bits.Len(uint(n)))

    return ans
}