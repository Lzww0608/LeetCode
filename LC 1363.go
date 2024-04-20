func largestMultipleOfThree(digits []int) string {
    sort.Slice(digits, func(i, j int) bool {
        return digits[i] > digits[j]
    })
    if digits[0] == 0 {
        return "0"
    }

    m := make([][]int, 3)
    sum := 0
    builder := strings.Builder{}

    for _, x := range digits {
        m[x%3] = append(m[x%3], x)
        sum += x
    }

    if sum % 3 == 0 {
        for _, x := range digits {
            builder.WriteByte(byte(x + '0'))
        }
        return builder.String()
    }

    if sum % 3 == 1 {
        if len(m[1]) > 0 {
            m[1] = m[1][:len(m[1])-1]
        } else {
            m[2] = m[2][:len(m[2])-2]
        }
    } else {
        if len(m[2]) > 0 {
            m[2] = m[2][:len(m[2])-1]
        } else {
            m[1] = m[1][:len(m[1])-2]
        }
    }

    ans := []int{}
    for i := range m {
        for _, x := range m[i] {
            ans = append(ans, x)
        }
    }

    sort.Slice(ans, func(i, j int) bool {
        return ans[i] > ans[j]
    })
    if len(ans) > 0 && ans[0] == 0 {
        return "0"
    }
    for _, x := range ans {
        builder.WriteByte(byte(x + '0'))
    }
    return builder.String()
}