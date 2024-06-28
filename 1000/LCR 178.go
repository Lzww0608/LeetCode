func trainingPlan(actions []int) (ans int) {
    for i := 1; i < (1 << 31); i <<= 1 {
        cnt := 0
        for _, x := range actions {
            if x & i != 0 {
                cnt++
            }
        }
        if cnt % 3 != 0 {
            ans = ans | i
        }
    }

    return
}
