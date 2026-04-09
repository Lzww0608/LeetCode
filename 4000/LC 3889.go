func mirrorFrequency(s string) (ans int) {
    cnt := [50]int{}
    
    for i := range s {
        if s[i] >= 'a' && s[i] <= 'z' {
            x := int(s[i] - 'a')
            cnt[x + 20]++
        } else {
            x := int(s[i] - '0')
            cnt[x]++
        }
    }

    for i := range s {
        if s[i] >= 'a' && s[i] <= 'z' {
            x := int(s[i] - 'a') + 20
            y := 25 - int(s[i] - 'a') + 20
            ans += abs(cnt[x] - cnt[y])
            cnt[x], cnt[y] = 0, 0
        } else {
            x := int(s[i] - '0')
            y := 9 - x 
            ans += abs(cnt[x] - cnt[y])
            cnt[x], cnt[y] = 0, 0
        }
    }

    return
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}