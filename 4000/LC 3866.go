func firstUniqueEven(nums []int) int {
    cnt := make(map[int]int)
    for _, x := range nums {
        if x & 1 == 0 {
            cnt[x]++
        }
    }

    for _, x := range nums {
        if x & 1 == 0 && cnt[x] == 1 {
            return x
        }
    }
    return -1
}