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

    for i := n; i >= 0; i-- {
        fmt.Println(fre[i])
        if len(fre[i]) > 0 {
            ans = append(ans, fre[i]...)
            k -= len(fre[i])
            if k <= 0 {
                return 
            }
        }
    }

    return 
}
