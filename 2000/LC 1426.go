func countElements(arr []int) (ans int) {
    cnt := make([]int, 1001)
    for _, x := range arr {
        cnt[x]++
    }

    for i := 0; i < len(cnt) - 1; i++ {
        if cnt[i+1] > 0 {
            ans += cnt[i];
        }
    }

    return 
}