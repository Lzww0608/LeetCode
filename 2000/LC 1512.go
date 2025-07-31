func numIdenticalPairs(nums []int) (ans int) {
    cnt := make(map[int]int)
    for _, x := range nums {
        ans += cnt[x]
        cnt[x]++
    }

    return 
}