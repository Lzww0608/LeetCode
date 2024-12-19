func sumOfNumberAndReverse(num int) bool {
    for i := num / 10; i <= num; i++ {
        if i + revese(i) == num {
            return true
        }
    }

    return false
}


func revese(n int) (ans int) {
    for n > 0 {
        ans = ans * 10 + n % 10
        n /= 10
    }
    return
}