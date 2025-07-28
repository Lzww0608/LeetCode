func maxTurbulenceSize(arr []int) (ans int) {
    n := len(arr)
    cur := 1
    ans = 1
    for i := 0; i < n - 1; i++ {
        if i & 1 == 1 && arr[i] > arr[i + 1] {
            cur++
        } else if i & 1 == 0 && arr[i] < arr[i + 1] {
            cur++
        } else {
            cur = 1
        }
        ans = max(ans, cur)
    }

    cur = 1
    for i := 0; i < n - 1; i++ {
        if i & 1 == 0 && arr[i] > arr[i + 1] {
            cur++
        } else if i & 1 == 1 && arr[i] < arr[i + 1] {
            cur++
        } else {
            cur = 1
        }
        ans = max(ans, cur)
    }

    return
}