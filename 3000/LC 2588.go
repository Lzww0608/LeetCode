const N int = 26
func beautifulSubarrays(nums []int) int64 {
    ans := 0

    mask := 0
    cnt := make(map[int]int)
    cnt[0]++
    for _, x := range nums {
        mask ^= x
        ans += cnt[mask]
        cnt[mask]++
    }

    return int64(ans)
}