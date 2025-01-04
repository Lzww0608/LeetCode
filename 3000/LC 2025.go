func waysToPartition(nums []int, k int) (ans int) {
    n := len(nums)
    sum := 0
    for _, x := range nums {
        sum += x 
    }

    pre := 0
    m1 := make(map[int]int)
    for _, x := range nums[:n-1] {
        pre += x
        suf := sum - pre
        m1[pre - suf]++
    }

    ans = m1[0]
    m2 := make(map[int]int)
    pre = 0
    //fmt.Println(m1)
    for _, x := range nums {
        d := k - x 
        pre += x
        suf := sum - pre
        cur := m2[d] + m1[-d]
        ans = max(ans, cur)
        m2[pre - suf]++
        m1[pre - suf]--
        //fmt.Println(ans)
    }

    return
}