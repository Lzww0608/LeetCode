func longestMountain(arr []int) int {
    n := len(arr)
    ans := 0
    l, r := 0, 0
    for r < n {
        for r < n - 1 && arr[r] < arr[r+1] {
            r++
        }
        if l == r {
            l, r = l + 1, r + 1
            continue
        } 
        t := r - l 
        for l = r; l < n - 1 && arr[l] > arr[l+1]; l++ {}
        if l > r {
            ans = max(ans, t + l - r + 1)
            r = l
        } else {
            l, r = l + 1, r + 1
        }
    }
    return ans
}