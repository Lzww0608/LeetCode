func minOperations(nums []int) (ans int) {
    cnt := make(map[int]int)
    for _, x := range nums {
        cnt[x]++
    }

    for _, x := range cnt {
        if x == 1 {
            return -1
        }
        ans += x / 3
        if x % 3 != 0 {
            ans++
        }
    }

    return
}