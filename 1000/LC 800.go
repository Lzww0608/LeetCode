
import "strconv"
func similarRGB(color string) string {
    builder := strings.Builder{}
    builder.WriteString("#")
    for i := 1; i < len(color); i += 2 {
        s := color[i:i+2]
        t, _ := strconv.ParseInt(s, 16, 32)
        x := t / 17
        if t % 17 > 8 {
            x++
        }
        c := byte('0' + x)
        if x >= 10 {
            c = byte('a' + x - 10)
        }
        builder.WriteByte(c)
        builder.WriteByte(c)
    }

    return builder.String()
}