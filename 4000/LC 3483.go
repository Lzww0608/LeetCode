func totalNumbers(digits []int) (ans int) {
    cnt := [10]int{}
    for _, x := range digits {
        cnt[x]++
    }
    for i := 0; i < 10; i += 2 {
        if cnt[i] == 0 {
            continue
        }
        cnt[i]--
        for j := range 10 {
            if cnt[j] == 0 {
                continue
            }
            cnt[j]--
            for k := 1; k < 10; k++ {
                if cnt[k] > 0 {
                    ans++
                }
            }
            cnt[j]++
        }
        cnt[i]++
    }

    return 
}