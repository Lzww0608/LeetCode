func findTheDistanceValue(arr1 []int, arr2 []int, d int) (ans int) {
    sort.Ints(arr1)
    sort.Ints(arr2)
    m, n := len(arr1), len(arr2)
    i, j := 0, 0

    for i < m {
        if abs(arr1[i] - arr2[j]) <= d {
            i++
        } else {
            if arr1[i] > arr2[j] && j < n - 1 {
                j++
            } else {
                ans++
                i++
            }
        }
    }

    return ans
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x 
}