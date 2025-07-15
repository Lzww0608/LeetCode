
import "math"
func maxCount(banned []int, n int, maxSum int64) (ans int) {
    banned = append(banned, 0)
    banned = append(banned, n + 1)
    sort.Ints(banned)
    j := 0
    
    for i := range banned {
        if i == len(banned) - 1 || banned[i] != banned[i+1] {
            banned[j] = banned[i]
            j++
        }
    }

    for i := 1; i < j; i++ {
        l := banned[i-1] + 1
        cnt := banned[i] - banned[i-1] - 1
        cur := cnt * l + cnt * (cnt - 1) / 2
        if cur < int(maxSum) {
            maxSum -= int64(cur)
            ans += cnt
        } else {
            ans += int((float64(1 - l * 2) + math.Sqrt(8 * float64(maxSum) + float64(1 - l * 2) * float64(1 - l * 2))) / 2)
            break
        }
    }

    return ans 
}