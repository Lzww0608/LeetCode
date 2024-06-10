func maxSatisfaction(satisfaction []int) int {
    sort.Sort(sort.Reverse(sort.IntSlice(satisfaction)))

    ans, pre := 0, 0
    for _, x := range satisfaction {
        pre += x
        if pre < 0 {
            break
        } 
        ans += pre
    }

    return ans
}
