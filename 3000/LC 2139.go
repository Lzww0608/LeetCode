func minMoves(target int, maxDoubles int) (ans int) {
    for target > 1 {
        if maxDoubles == 0 {
            ans += target - 1
            break
        } else if target & 1 == 1 {
            target--
            ans++
        } else {
            target >>= 1
            maxDoubles--
            ans++
        }
    }

    return
}