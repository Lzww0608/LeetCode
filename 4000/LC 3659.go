func partitionArray(nums []int, k int) bool {
    if len(nums) % k != 0 {
        return false
    }
    cnt := make(map[int]int)
    for _, x := range nums {
        if cnt[x]++; cnt[x] > len(nums) / k {
            return false
        }
    }

    return true
}