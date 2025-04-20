func findSmallestInteger(nums []int, k int) int {
    cnt := make([]int, k)
    for _, x := range nums {
        cnt[(x % k + k) % k]++
    }

    for i := 0; ; i++ {
        if cnt[i % k] == 0 {
            return i
        }
        cnt[i % k]--
    }

    return 0
}