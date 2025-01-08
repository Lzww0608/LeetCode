
import "strings"
func largestGoodInteger(num string) string {
    ans := -1
    for i := 0; i + 2 < len(num); i++ {
        if num[i] == num[i+1] && num[i+1] == num[i+2] {
            x := int(num[i] - '0')
            ans = max(ans, x)
        } 
    }

    if ans == -1 {
        return ""
    }

    return strings.Repeat(string('0' + ans), 3)
}