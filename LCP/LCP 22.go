func paintingPlan(n int, k int) (ans int) {
    if k == 0 || k == n * n {
        return 1
    }
    for r := 0; r < (1 << n); r++ {
        for c := 0; c < (1 << n); c++ {
            cnt := bits.OnesCount(uint(r))
            sum := cnt * n
            for i := range n {
                if c & (1 << i) != 0 {
                    sum += n - cnt
                }
            }

            if sum == k {
                ans++
            }
        }
    }

    return
}