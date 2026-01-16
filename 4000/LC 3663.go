func getLeastFrequentDigit(n int) int {
    cnt := [10]int{}
    
    ans := 10
    for n > 0 {
        x := n % 10 
        n /= 10
        cnt[x]++
        ans = min(ans, x)
    }

    for i := range cnt {
        if cnt[i] > 0 && cnt[i] < cnt[ans] {
            ans = i
        }
    }

    return ans
}