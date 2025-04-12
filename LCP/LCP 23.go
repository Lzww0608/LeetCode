func isMagic(target []int) bool {
    n := len(target)
    a := make([]int, 0, n)
    for i := 2; i <= n; i += 2 {
        a = append(a, i)
    }
    for i := 1; i <= n; i += 2 {
        a = append(a, i)
    }

    //fmt.Println(a)
    k := -1
    for len(a) > 0 {
        i := 0
        for i < len(a) {
            if a[i] == target[i] {
                i++
            } else {
                break
            }
        }

        if i == len(a) {
            break
        }

        if i == 0 || k != -1 && k != i {
            return false
        }
        k = i
        a = a[i:]
        target = target[i:]
        b := make([]int, 0, len(a))
        for i := 1; i < len(a); i += 2 {
            b = append(b, a[i])
        }
        for i := 0; i < len(a); i += 2 {
            b = append(b, a[i])
        }
        a = b 
    }

    return true
}