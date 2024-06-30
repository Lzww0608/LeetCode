func multiply(num1 string, num2 string) string {
    if num1[0] == '0' || num2[0] == '0' {
        return "0"
    }
    prod := make([]int, 400)
    n, m := len(num1), len(num2)
    for i := n - 1; i >= 0; i-- {
        x := int(num1[i] - '0')
        add, k := 0, n - 1 - i
        for j := m - 1; j >= 0; j-- {
            y := int(num2[j] - '0')
            t := add + x * y
            prod[k] += t
            add = prod[k] / 10
            prod[k] %= 10
            k++
        }
        prod[k] = add 
    }
    ans := strings.Builder{}
    i := len(prod) - 1
    for prod[i] == 0 {
        i--
    }
    for i >= 0 {
        ans.WriteString(strconv.Itoa(prod[i]))
        i--
    }
    return ans.String()
}
