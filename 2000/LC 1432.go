func maxDiff(num int) int {
    tmp := strconv.Itoa(num)
    a := []byte(tmp)
    b := []byte(tmp)
    n := len(a)
    i := 0
    for i < n && a[i] == '9' {
        i++
    }
    for j := n - 1; j >= i; j-- {
        if a[j] == a[i] {
            a[j] = '9'
        }
    }

    if b[0] == '1' {
        i = 0
        for i < n && (b[i] == '1' || b[i] == '0') {
            i++
        }
        for j := n - 1; j >= i; j-- {
            if b[j] == b[i] {
                b[j] = '0'
            }
        }
    } else {
        for j := n - 1; j >= 0; j-- {
            if b[j] == b[0] {
                b[j] = '1'
            }
        }
    }

    x, _ := strconv.Atoi(string(a))
    y, _ := strconv.Atoi(string(b))
    //fmt.Println(x, y)
    return x - y
}