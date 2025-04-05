
import "math"
func sumGame(num string) bool {
    n := len(num)
    var sum float64 = 0
    for i := 0; i < n; i++ {
        if num[i] == '?' {
            if i < n / 2 {
                sum += 4.5
            } else {
                sum -= 4.5
            }
        } else {
            if i < n / 2 {
                sum += float64(int(num[i] - '0'))
            } else {
                sum -= float64(int(num[i] - '0'))
            }
        }
    }

    return math.Abs(sum) > 1e-5
}