func fractionAddition(s string) string {
    a := []int{}
    b := []int{}
    dom := 1
    
    n := len(s)
    for i := 0; i < n; i++ {
        neg := 1
        if s[i] == '-' {
            neg = -1
            i++
        } else if s[i] == '+' {
            i++
        }

        x := int(s[i] - '0')
        i++
        if s[i] != '/' {
            x = x * 10 + int(s[i] - '0')
            i++
        }
        a = append(a, neg * x)
        i++
        y := int(s[i] - '0')
        if i + 1 < n && s[i + 1] != '+' && s[i + 1] != '-' {
            i++
            y = y * 10 + int(s[i] - '0')
        }
        b = append(b, y)
        dom *= b[len(b)-1]
    }

    sum := 0
    for i, x := range a {
        sum += x * (dom / b[i])
    }

    t := abs(sum)
    g := gcd(t, dom)
    t /= g 
    dom /= g

    ans := strconv.Itoa(t) + "/" + strconv.Itoa(dom)
    if sum < 0 {
        return "-" + ans
    }
    return ans
}

func gcd(a, b int) int {
    for b != 0 {
        a, b = b, a % b
    }

    return a
}

func abs(x int) int {
    if x < 0 {
        return -x
    }

    return x
}