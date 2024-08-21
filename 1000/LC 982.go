func countTriplets(nums []int) (ans int) {
    cnt := map[int]int{}
    cnt[0] = len(nums)
    for _, m := range nums {
        m ^= 0xffff
        for s := m; s > 0; s = (s - 1) & m {
            cnt[s]++
        }
    }

    for _, x := range nums {
        for _, y := range nums {
            ans += cnt[x & y]
        }
    }

    return ans
}