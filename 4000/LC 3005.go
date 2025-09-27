func maxFrequencyElements(nums []int) (ans int) {
    cnt := make(map[int]int)
    mx := 0

    for _, x := range nums {
        cnt[x]++
        if cnt[x] > mx {
            mx = cnt[x]
            ans = mx
        } else if cnt[x] == mx {
            ans += mx
        }
    }

    return 
}