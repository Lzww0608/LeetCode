
import (
	"strconv"
	"strings"
)
func fractionToDecimal(a int, b int) string {
    if a % b == 0 {
        return strconv.Itoa(a / b)
    }

    ans := strings.Builder{}
    if a * b < 0 {
        ans.WriteByte('-')
    }
    a, b = abs(a), abs(b)
    ans.WriteString(strconv.Itoa(a / b))
    ans.WriteByte('.')

    a %= b
    m := make(map[int]int) 
    for a != 0 {
        m[a] = ans.Len()
        a *= 10
        ans.WriteString(strconv.Itoa(a / b))
        a %= b 
        if v, ok := m[a]; ok {
            s1 := ans.String()[:v]
            s2 := ans.String()[v:]
            return s1 + "(" + s2 + ")"
        }
    }

    return ans.String()
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}