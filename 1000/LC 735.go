func asteroidCollision(a []int) []int {
    ans := []int{}
    for _, x := range a {
        f := true
        for f && x < 0 && len(ans) > 0 && ans[len(ans)-1] > 0 {
            if ans[len(ans)-1] == abs(x) {
                ans = ans[:len(ans)-1]
                f = false
            } else if ans[len(ans)-1] > abs(x) {
                f = false
            } else {
                ans = ans[:len(ans)-1]
            }
        }
        if f {
            ans = append(ans, x)
        }
    }
    return ans
}

func abs(x int) int {
    if x < 0 {
        return -x 
    }
    return x
}
