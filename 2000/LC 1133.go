const N int = 1_001
func largestUniqueNumber(nums []int) int {
    cnt := make([]int, N)
    for _, x := range nums {
        cnt[x]++
    }

    for i := N - 1; i >= 0; i-- {
        if cnt[i] == 1 {
            return i
        }
    }

    return -1
}