func subarraysDivByK(nums []int, k int) (ans int) {
    cnt := make(map[int]int)

    sum := 0
    cnt[0] = 1
    for _, x := range nums {
        sum = (sum + x) % k
        if sum < 0 {
            sum += k
        }
        ans += cnt[sum]
        cnt[sum]++
    }

    return 
}