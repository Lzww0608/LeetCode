func findMaxK(nums []int) int {
    m := map[int]bool{}
    ans := -1
    for _, x := range nums {
        m[x] = true 
        if m[-x] {
            ans = max(ans, x, -x)
        }
    }

    return ans
}