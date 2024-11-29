
import (
	"strconv"
)
func kMirror(k int, n int) int64 {
    ans, num := 0, 0

    for n > 0 {
        num = next(num)
        s := strconv.FormatInt(int64(num), k)
        if check(s) {
            ans += num 
            n--
        }
    }

    return int64(ans)
}

func next(x int) int {
    s := []byte(strconv.Itoa(x))
    n := len(s)
    for i := n / 2; i >= 0; i-- {
        if s[i] != '9' {
            s[i]++
            if n - 1 - i != i {
                s[n-1-i]++
            }
            for j := i + 1; j <= n / 2; j++ {
                s[j] = '0'
                s[n-1-j] = '0'
            }
            ans, _ := strconv.Atoi(string(s))
            return ans
        }
    }

    ans := 1
    for i := 0; i < n; i++ {
        ans *= 10
    }
    ans++
    return ans
}

func check(s string) bool {
    n := len(s)
    for i, j := 0, n - 1; i < j; i, j = i + 1, j - 1 {
        if s[i] != s[j] {
            return false
        }
    }

    return true
}