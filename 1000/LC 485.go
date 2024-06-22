func findMaxConsecutiveOnes(nums []int) (ans int) {
    cnt := 0
    for _, x := range nums {
        if x == 1 {
            cnt++
        } else {
            ans = max(ans, cnt)
            cnt = 0
        }
    }
    ans = max(ans, cnt)
    return
}
