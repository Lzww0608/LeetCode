func badSensor(a []int, b []int) int {
    n := len(a)
    for i := range n - 1 {
        if a[i] != b[i] {
            f1 := check(a[i + 1:], b[i: n - 1])
            f2 := check(a[i: n - 1], b[i + 1:])
            if f1 && !f2 {
                return 2
            } else if f2 && !f1 {
                return 1
            }
            return -1
        }
    }

    return -1
}

func check(a, b []int) bool {
    for i := range a {
        if a[i] != b[i] {
            return false
        }
    }

    return true
}