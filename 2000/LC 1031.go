func maxSumTwoNoOverlap(nums []int, firstLen int, secondLen int) (ans int) {
    n := len(nums)
    prefix := make([]int, n + 1)
    for i, x := range nums {
        prefix[i+1] = prefix[i] + x
    }

    sum_f, sum_s := 0, 0
    for i := firstLen + secondLen; i <= n; i++ {
        sum_f = max(sum_f, prefix[i-secondLen] - prefix[i-secondLen-firstLen])
        sum_s = max(sum_s, prefix[i-firstLen] - prefix[i-firstLen-secondLen])
        ans = max(ans, sum_f + prefix[i] - prefix[i-secondLen], sum_s + prefix[i] - prefix[i-firstLen])
    }

    return 
}
