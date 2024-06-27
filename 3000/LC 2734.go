func smallestString(str string) string {
    s := []byte(str)
    n := len(s)
    i, f := 0, true

    for i < n && f {
        j := i
        for j < n && s[j] != 'a' {
            j++
        }
        if j > i {
            for ;i < j; i++ {
                s[i]--
            }
            f = false
        } else if (i == j && j == n - 1) {
            f = false
            s[i] = 'z'
        } else {
            i = j + 1
        }
    }


    return string(s)
}