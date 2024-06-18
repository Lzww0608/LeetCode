func canReach(s string, minJump int, maxJump int) (ans bool) {
    n := len(s)
    pre := make([]int, n + 1)
    pre[1] = 1
    for i := 1; i < n; i++ {
        ans = false
        if s[i] == '1' || i < minJump {
            pre[i+1] = pre[i]
        } else {
            l, r := max(0, i - maxJump), max(0, i - minJump) + 1
            if pre[r] > pre[l] {
                ans = true
                pre[i+1] = pre[i] + 1
            } else {
                pre[i+1] = pre[i]
            }
        }
    }

    return
}