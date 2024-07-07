func masterMind(solution string, guess string) (ans []int) {
    hit, pseudo_hit := 0, 0
    m := map[byte]int{}

    for i := range solution {
        if solution[i] == guess[i] {
            hit++
        } else {
            if m[guess[i]] < 0 {
                pseudo_hit++
            }
            m[guess[i]] += 1
            
            if m[solution[i]] > 0 {
                pseudo_hit++
            }
            m[solution[i]]--
        }
    }

    ans = []int{hit, pseudo_hit}
    return
}