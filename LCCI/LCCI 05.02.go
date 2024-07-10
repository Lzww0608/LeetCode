const EPS float64 = 0.0000001
func printBin(num float64) string {
    builder := strings.Builder{}
    builder.WriteString("0.")
    for builder.Len() <= 32  {
        num *= 2.0
        if math.Abs(num - 1) < EPS {
            builder.WriteString("1")
            break
        } else if num - 1 > EPS {
            builder.WriteString("1")
            num -= 1.0
        } else {
            builder.WriteString("0")
        }
    }

    if builder.Len() > 32 {
        return "ERROR"
    }

    return builder.String()
}