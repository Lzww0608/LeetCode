func isPossibleDivide(nums []int, k int) bool {
    n := len(nums)
    if n % k != 0 {
        return false
    }
    sort.Ints(nums)
    m := map[int]int{}
    for _, x := range nums {
        m[x]++
    }
    for _, x := range nums {
        if m[x] == 0 {
            continue
        }
        for j := 0; j < k; j++ {
            if m[x+j] == 0 {
                return false
            }
            m[x+j]--
        }
    }
    return true
}