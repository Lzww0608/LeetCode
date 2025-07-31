
import "slices"
func kthLuckyNumber(k int) string {
    k++
    ans := []byte{}
    for k > 1 {
        if k & 1 == 0 {
            ans = append(ans, '4')
        } else {
            ans = append(ans, '7')
        }

        k >>= 1
    }

    slices.Reverse(ans)

    return string(ans)
}