func twoEditWords(queries []string, dictionary []string) (ans []string) {
    n := len(queries[0])
    check := func(i, j int) bool {
        cnt := 0
        for k := 0; k < n; k++ {
            if queries[i][k] != dictionary[j][k] {
                cnt++
            }
        }

        return cnt <= 2
    }    

    for i := range queries {
        for j := range dictionary {
            if check(i, j) {
                ans = append(ans, queries[i])
                break
            }
        }
    }

    return 
}

