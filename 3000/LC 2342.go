func maximumSum(nums []int) int {
    
    getSum := func(x int) int {
        res := 0
        for x > 0 {
            res += x % 10
            x /= 10
        }
        return res
    }

    m := map[int]int{}
    ans := -1

    for _, x := range nums {
        t := getSum(x)
        if v, ok := m[t]; ok {
            ans = max(ans, v + x)
        }
        m[t] = max(m[t], x)
    }
    
    return ans 
}
