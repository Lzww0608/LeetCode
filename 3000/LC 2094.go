func findEvenNumbers(digits []int) (ans []int) {
    cnt := [10]int{}
    for _, x := range digits {
        cnt[x]++
    }

    for i := 1; i < 10; i++ {
        if cnt[i] == 0 {
            continue
        }
        cnt[i]--
        for j := 0; j < 10; j++ {
            if cnt[j] == 0 {
                continue
            }
            cnt[j]--
            for k := 0; k < 10; k += 2 {
                if cnt[k] > 0 {
                    ans = append(ans, i * 100 + j * 10 + k)
                }
            }
            cnt[j]++
        }
        cnt[i]++
    }

    return
}