func getConcatenation(nums []int) []int {
    ans := nums 
    ans = append(ans, nums...)
    return ans
}