
import "reflect"
func minOperations(a []int, b []int, k int) int64 {
    if k == 0 {
        if reflect.DeepEqual(a, b) {
            return 0
        }
        return -1
    }
    cnt := 0
    ans := 0
    n := len(a)
    for i := 0; i < n; i++ {
        if abs(a[i] - b[i]) % k != 0 {
            return -1
        }
        cnt += (a[i] - b[i]) / k
        ans += abs(a[i] - b[i]) / k
    }

    if cnt != 0 {
        return -1
    }

    return int64(ans) / 2
}

func abs(x int) int {
    if x < 0 {
        return -x
    }

    return x
}