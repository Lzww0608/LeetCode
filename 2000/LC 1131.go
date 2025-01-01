
import "math"
func maxAbsValExpr(arr1 []int, arr2 []int) int {
    // j >= i 
    // j - i + arr1[i] - arr1[j] + arr2[i] - arr2[j]
    // (arr1[i] + arr2[i] - i) + (-arr1[j] - arr2[j] + j)
    // (-arr1[i] + arr2[i] - i) + (arr1[j] - arr2[j] + j)
    // (arr1[i] - arr2[i] - i) + (-arr1[j] + arr2[j] + j)
    // (-arr1[i] - arr2[i] - i) + (arr1[j] + arr2[j] + j)
    a := [2]int{math.MinInt32, math.MinInt32}
    b := [2]int{math.MinInt32, math.MinInt32}
    c := [2]int{math.MinInt32, math.MinInt32}
    d := [2]int{math.MinInt32, math.MinInt32}

    for i := range arr1 {
        a[0] = max(a[0], arr1[i] + arr2[i] - i)
        a[1] = max(a[1], -arr1[i] - arr2[i] + i)
        b[0] = max(b[0], -arr1[i] + arr2[i] - i)
        b[1] = max(b[1], arr1[i] - arr2[i] + i)
        c[0] = max(c[0], arr1[i] - arr2[i] - i)
        c[1] = max(c[1], -arr1[i] + arr2[i] + i)
        d[0] = max(d[0], -arr1[i] - arr2[i] - i)
        d[1] = max(d[1], arr1[i] + arr2[i] + i)
    }

    return max(a[0] + a[1], b[0] + b[1], c[0] + c[1], d[0] + d[1])
}