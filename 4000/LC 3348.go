func smallestNumber(num string, t int64) string {
    tmp := int(t)
    for i := 9; i > 1; i-- {
        for tmp % i == 0 {
            tmp /= i
        }
    }

    if tmp > 1 {
        return "-1"
    }

    n := len(num)
    a := make([]int, n + 1)
    a[0] = int(t)
    j := n - 1
    for i, c := range num {
        if c == '0' {
            j = i 
            break
        }
        a[i + 1] = a[i] / gcd(a[i], int(c - '0'))
    }
    if a[n] == 1 {
        return num
    }

    s := []byte(num)
    for i := j; i >= 0; i-- {
        for s[i] += 1; s[i] <= '9'; s[i]++ {
            tmp = a[i] / gcd(a[i], int(s[i] - '0'))
            k := 9
            for j := n - 1; j > i; j-- {
                for tmp % k > 0 {
                    k--
                }
                tmp /= k 
                s[j] = '0' + byte(k)
            }

            if tmp == 1 {
                return string(s)
            }
        }
    }

    ans := make([]byte, 0, n + 1)
    for i := int64(9); i > 1; i-- {
        for t % i == 0 {
            ans = append(ans, byte('0' + i))
            t /= i
        }
    }

    for len(ans) <= n {
        ans = append(ans, '1')
    }
    slices.Reverse(ans)
    return string(ans)
}

func gcd(a, b int) int {
    for b != 0 {
        a, b = b, a % b
    }

    return a
}