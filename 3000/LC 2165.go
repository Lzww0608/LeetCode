func smallestNumber(num int64) int64 {
    cnt := make([]int, 10)
    for t := abs(num); t != 0; t /= 10 {
        x := t % 10
        cnt[x]++
    }

    ans := 0
    if num >= 0 {
        for i := 1; i < 10; i++ {
            if cnt[i] == 0 {
                continue
            }
            for cnt[i] > 0 {
                ans = ans * 10 + i
                for cnt[0] > 0 {
                    cnt[0]--
                    ans = ans * 10
                }
                cnt[i]--
            }
        }
        return int64(ans)
    }
    for i := 9; i >= 0; i-- {
        for cnt[i] > 0 {
            cnt[i]--
            ans = ans * 10 + i
        }
    }

    return int64(-ans)
}

func abs(x int64) int64 {
    if x < 0 {
        return -x
    }

    return x
}