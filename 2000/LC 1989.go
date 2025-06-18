func catchMaximumAmountofPeople(team []int, dist int) (ans int) {
    //n := len(team)
    a := []int{}
    for i, x := range team {
        if x == 0 {
            a = append(a, i)
        }
    }

    for i, x := range team {
        if len(a) == 0 {
            break
        }
        if x == 1 {
            l, r := i - dist, i + dist
            for len(a) > 0 && a[0] < l {
                a = a[1:]
            }
            if len(a) > 0 && a[0] <= r {
                a = a[1:]
                ans++
            }
        }
    }

    return
}