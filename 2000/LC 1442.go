func countTriplets(arr []int) int {
    n := len(arr)
    prefix := make([]int, n+1)
    for i, x := range arr {
        prefix[i+1] = prefix[i] ^ x
    }
    ans := 0
    for i := 0; i <= n; i++ {
        for j := i+1; j < n; j++ {
            if  prefix[i] ==  prefix[j+1] {
                ans += j - i
            }
        }
    }
    return ans

}
