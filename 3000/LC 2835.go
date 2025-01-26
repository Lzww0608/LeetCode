func minOperations(nums []int, target int) (ans int) {
    cnt := make([]int, 32)
    for _, x := range nums {
        cnt[bits.Len(uint(x))-1]++
    }
    for i := 0; (1 << i) <= target; i++ {
        mask := 1 << i 
        if target & mask == 0 {
            if i + 1 < len(cnt) {
                cnt[i+1] += cnt[i] / 2
            }
            continue
        }
        j := i
        for j < len(cnt) && cnt[j] == 0 {
            j++
        }
        if j == len(cnt) {
            return -1
        }
        if j > i {
            cnt[j]--
        }
        ans += j - i
        for t := j - 1; t >= i; t-- {
            cnt[t]++
        }
        if i + 1 < len(cnt) {
            cnt[i+1] += (cnt[i] - 1) / 2
        }
    }

    return
}