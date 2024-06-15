func addNegabinary(a []int, b []int) (ans []int) {
    i, j := len(a) - 1, len(b) - 1
    for add := 0; add != 0 || i >= 0 || j >= 0; i, j = i - 1, j - 1 {
        x, y := 0, 0
        if i >= 0 {
            x = a[i] 
        }
        if j >= 0 {
            y = b[j]
        }
        tmp := x + y + add
        add = 0
        if tmp >= 2 {
            tmp -= 2
            add--
        } else if tmp == -1 {
            tmp = 1
            add++
        }
        
        ans = append(ans, tmp)
    }

    for len(ans) > 1 && ans[len(ans)-1] == 0 {
        ans = ans[:len(ans)-1]
    }

    for i := 0; i * 2 < len(ans); i++ {
        ans[i], ans[len(ans)-1-i] = ans[len(ans)-1-i], ans[i]
    }

    return 
}