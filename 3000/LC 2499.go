func minimumTotalCost(a []int, b []int) int64 {
    n := len(a)
    cnts := make(map[int]int)
    cnt, num, swap := 0, 0, 0
    ans := 0
    for i := 0; i < n; i++ {
        if a[i] == b[i] {
            x := a[i]
            swap++
            cnts[x]++
            ans += i
            if cnts[x] > cnts[num] {
                num = x
            }
            cnt = cnts[num]
        }
    }

    for i := 0; i < n; i++ {
        if cnt * 2 <= swap {
            break
        }
        if a[i] != b[i] && a[i] != num && b[i] != num {
            swap++
            ans += i 
        }
    }

    if cnt * 2 > swap {
        return -1
    }

    return int64(ans)
}