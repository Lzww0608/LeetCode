func constructArray(n int, k int) []int {
    ans := make([]int, n)
    l, r := 1, n 
    for i := 0; i < k; i++ {
        if i & 1 == 0 {
            ans[i] = l
            l++
        } else {
            ans[i] = r 
            r--
        }
    }

    for i := k; i < n; i++ {
        if k & 1 == 0 {
            ans[i] = r 
            r--
        } else {
            ans[i] = l 
            l++
        }
    }

    return ans
}