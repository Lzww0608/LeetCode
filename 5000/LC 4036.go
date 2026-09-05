func largestString(nums []int) []string {
    n := len(nums)
    ans := make([]string, n)

    for i, x := range nums {
        cur := []byte{}
        for j := 25; j >= 0; j-- {
            y := 1 << j 
            for x >= y {
                x -= y
                cur = append(cur, byte('a' + j))
            }
        }
        ans[i] = string(cur)
    }

    return ans
}