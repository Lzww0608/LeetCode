func numberOfPairs(nums []int) []int {
    cnt := make(map[int]int)
    for _, x := range nums {
        cnt[x]++
    }

    ans := make([]int, 2)
    for _, v := range cnt {
        ans[0] += v / 2 
        ans[1] += v & 1
    }

    return ans
}