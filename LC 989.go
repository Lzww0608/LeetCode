func addToArrayForm(num []int, k int) []int {
    n := len(num)
    ans := []int{}
    i, add := n - 1, 0
    for i >= 0 || add > 0 || k > 0 {
        t := add + k % 10
        if i >= 0 {
            t += num[i]
            i--
        }
        k /= 10
        add, t = t / 10, t % 10 
        ans = append(ans, t)    
    }
    for j, k := 0, len(ans) - 1; j < k; j, k = j + 1, k - 1 {
        ans[j], ans[k] = ans[k], ans[j]
    }

    return ans
}