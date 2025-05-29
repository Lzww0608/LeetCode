func differByOne(dict []string) bool {
    n := len(dict)
    for i := 0; i < n; i++ {
        for j := i + 1; j < n; j++ {
            if dif(dict[i], dict[j]) {
                return true
            }
        }
    }

    return false
}

func dif(a, b string) bool {
    cnt := 0
    for i := range a {
        if a[i] != b[i] {
            cnt++
            if cnt > 1 {
                return false
            }
        }
    }

    return cnt == 1
}