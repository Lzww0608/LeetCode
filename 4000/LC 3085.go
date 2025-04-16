
import "math"
const N int = 26
func minimumDeletions(word string, k int) (ans int) {
    m := make([]int, N)
    for i := range word {
        x := int(word[i] - 'a')
        m[x]++
    }

    ans = math.MaxInt
    for i := 0; i < N; i++ {
        if m[i] == 0 {
            continue
        }
        cnt := 0
        for j := 0; j < N; j++ {
            if m[j] < m[i] {
                cnt += m[j]
            } else if m[j] > m[i] + k {
                cnt += m[j] - m[i] - k
            }
        }        
        ans = min(ans, cnt)
    }

    return 
}