func trainingPlan(actions []int) []int {
    i, k := 0, len(actions)
    for i < k {
        if actions[i] & 1 == 1 {
            i++
        } else {
            k--
            actions[k], actions[i] = actions[i], actions[k] 
        }
    }

    return actions
}
