func complexNumberMultiply(num1 string, num2 string) string {
    num1 = num1[:len(num1) - 1]
    num2 = num2[:len(num2) - 1]
    s := strings.Split(num1, "+")
    t := strings.Split(num2, "+")
    a, _ := strconv.Atoi(s[0])
    b, _ := strconv.Atoi(s[1])
    c, _ := strconv.Atoi(t[0])
    d, _ := strconv.Atoi(t[1])

    x := a * c - b * d 
    y := a * d + b * c 
    return strconv.Itoa(x) + "+" + strconv.Itoa(y) + "i"
}