func numOfSubarrays(arr []int, k int, threshold int) (ans int) {
    target := k * threshold
    l, r := 0, 0
    for i := 0; i < k; i++ {
        r += arr[i]
    }
    if r >= target {
        ans++
    }

    for i := k; i < len(arr); i++ {
        r += arr[i]
        l += arr[i-k]
        if r - l >= target {
            ans++
        }
    }

    return 
}