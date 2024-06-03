func checkPalindromeFormation(a string, b string) bool {
    n := len(a)
    i, j := 0, n - 1
    
    checkA := func(i, j int) bool {
        for i < j {
            if a[i] != a[j] {
                return false
            }
            i++
            j--
        }
        return true
    }

    checkB := func(i, j int) bool {
        for i < j {
            if b[i] != b[j] {
                return false
            }
            i++
            j--
        }
        return true
    }

    for i < j {
        if a[i] != b[j] {
            if checkA(i, j) || checkB(i, j) {
                return true
            }
            break
        }
        i++
        j--
    }

    if i >= j {
        return true
    }

    for i, j = 0, n - 1; i < j; i, j = i + 1, j - 1 {
        if b[i] != a[j] {
            if checkA(i, j) || checkB(i, j) {
                return true
            }
            break
        }
    } 

    return i >= j
}
