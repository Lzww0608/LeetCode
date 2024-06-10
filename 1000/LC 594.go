func findLHS(nums []int) (ans int) {
    //sort.Ints(nums)

    m := map[int]int{}
    for _, x := range nums {
        m[x]++
        if m[x-1] > 0 {
            ans = max(ans, m[x] + m[x-1])
        } 
        if m[x+1] > 0 {
            ans = max(ans, m[x] + m[x+1])
        }
    }

    return 
}
