func distance(nums []int) []int64 {
    m := make(map[int][]int)
    for i, x := range nums {
        m[x] = append(m[x], i) 
    }

    n := len(nums)
    ans := make([]int64, n)
    for _, v := range m {
        sum := 0
        for _, x := range v {
            sum += x  
        }

        cur := 0
        sz := len(v)
        for j, x := range v {
            ans[x] = int64(sum - (sz - j) * x + j * x - cur)
            cur += x
            sum -= x
        }
    }
    return ans
}

