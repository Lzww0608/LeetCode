func toggleLightBulbs(bulbs []int) []int {
    m := make(map[int]int)
    for _, x := range bulbs {
        m[x] ^= 1
    }

    var ans []int 
    for k, v := range m {
        if v == 1 {
            ans = append(ans, k)
        }
    }
    sort.Ints(ans)

    return ans
}