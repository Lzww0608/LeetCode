const N int = 1000
func miceAndCheese(reward1 []int, reward2 []int, k int) (ans int) {
    dif := make([]int, 2 * N)
    for i := range reward1 {
        t := 1000 + reward1[i] - reward2[i]
        dif[t]++
        ans += reward2[i]
    }

    for i := len(dif) - 1; i > 0 && k > 0; i-- {
        for k > 0 && dif[i] > 0 {
            ans += i - 1000
            k--
            dif[i]--
        } 
    }
    
    return
}
