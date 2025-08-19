func minNumberOfFrogs(croakOfFrogs string) (ans int) {
    cnt := [4]int{}

    for _, c := range croakOfFrogs {
        if c == 'c' {
            cnt[0]++
        } else if c == 'r' {
            cnt[0]--
            cnt[1]++
        } else if c == 'o' {
            cnt[1]--
            cnt[2]++
        } else if c == 'a' {
            cnt[2]--
            cnt[3]++
        } else if c == 'k' {
            cnt[3]--
        }

        sum := 0
        for _, x := range cnt {
            if x < 0 {
                return -1
            }
            sum += x
        }
        ans = max(ans, sum)
    }

    for _, x := range cnt {
        if x != 0 {
            return -1
        }
    }

    return 
}