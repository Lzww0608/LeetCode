func numberCount(a int, b int) (ans int) {
    for x := a; x <= b; x++ {
        if check(x) {
            ans++
        }
    }

    return
}

func check(x int) bool {
    cnt := [10]bool{}
    for x > 0 {
        if cnt[x % 10] {
            return false
        }
        cnt[x % 10] = true 
        x /= 10
    }

    return true
}