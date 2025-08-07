const N int = 10
func minimumSum(num int) (ans int) {
    cnt := [N]int{}
    for num > 0 {
        cnt[num % 10]++
        num /= 10
    }

    t := 2
    for i := 0; i < N; i++ {
        for cnt[i] > 0 {
            if t > 0 {
                t--
                ans += i * 10
            } else {
                ans += i
            }
            cnt[i]--
        }
    }

    return
}