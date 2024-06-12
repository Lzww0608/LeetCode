func numSteps(s string) (ans int) {
    n, add := len(s), 0
    j := 0
    for j < n && s[j] == '0' {
        j++
    }
    for i := n - 1; i > j; i-- {
        x := int(s[i] - '0') + add
        add = x / 2
        if x & 1 == 0 {
            ans++ 
        } else {
            ans += 2
            add++
        }
    }

    if add > 0 {
        ans++
    }

    return
}