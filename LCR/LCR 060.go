func topKFrequent(nums []int, k int) (ans []int) {
    n := len(nums)
    fre := make([][]int, n + 1)
    m := map[int]int{}
    for _, x := range nums {
        m[x]++
    }
    for k, v := range m {
        fre[v] = append(fre[v], k)
    }

    for i := n; i > 0 && k > 0; i-- {
        if len(fre[i]) > 0 {
            for _, x := range fre[i] {
                ans = append(ans, x)
            }
            k -= len(fre[i])
        }
    }

    return
}
