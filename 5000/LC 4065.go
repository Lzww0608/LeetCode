const N int = 101
func rearrangeArray(nums []int) []int {
    cnt := make([]int, N)
    for _, x := range nums {
        cnt[x]++
    }

    ans := make([]int, 0, len(nums))
    for len(ans) < len(nums) {
        for i, x := range cnt {
            if x == 0 {
                continue
            }
            ans = append(ans, i)
            cnt[i]--
        }
    }

    return ans
}