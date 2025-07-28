import "slices"
func pancakeSort(arr []int) []int {
    if slices.IsSorted(arr) {
        return nil
    }

    n := len(arr)
    ans := make([]int, 0, n * 2)
    for i := 0; i < n; i++ {
        mx := 0
        for j := 1; j < n - i; j++ {
            if arr[mx] < arr[j] {
                mx = j
            }
        }
        ans = append(ans, mx + 1)
        ans = append(ans, n - i)
        slices.Reverse(arr[:mx + 1])
        slices.Reverse(arr[:n - i])
    }

    return ans
}