func maxFrequency(nums []int, k int, numOperations int) (ans int) {
    cnt := make(map[int]int)
    diff := make(map[int]int)

    for _, x := range nums {
        cnt[x]++
        diff[x] += 0
        diff[x-k]++
        diff[x+k+1]--
    }

    sum := 0
    for _, x := range slices.Sorted(maps.Keys(diff)) {
        sum += diff[x]
        ans = max(ans, cnt[x] + min(sum - cnt[x], numOperations))
    }

    return
}