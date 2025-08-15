func divideArray(nums []int) bool {
    cnt := make(map[int]int)
    for _, x := range nums {
        cnt[x]++
    }

    for _, x := range cnt {
        if x & 1 == 1 {
            return false
        }
    }

    return true
}