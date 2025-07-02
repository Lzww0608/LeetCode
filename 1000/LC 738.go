
import "strconv"
func monotoneIncreasingDigits(n int) (ans int) {
    s := []byte(strconv.Itoa(n))
    m := len(s)
    
    for i := 1; i < m; i++ {
        if s[i] < s[i-1] {
            s[i-1]--
            for i - 2 >= 0 && s[i-2] > s[i-1] {
                s[i-2]--
                s[i-1] = '9'
                i--
            }
            for j := i; j < m; j++ {
                s[j] = '9'
            }
        }
    } 
    ans, _ = strconv.Atoi(string(s))
    return
}