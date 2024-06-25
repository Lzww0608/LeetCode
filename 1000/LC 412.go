func fizzBuzz(n int) []string {
    ans := make([]string, n)
    for i := range ans {
        x := i + 1
        if x % 3 == 0 && x % 5 == 0 {
            ans[i] = "FizzBuzz"
        } else if x % 3 == 0 {
            ans[i] = "Fizz"
        } else if x % 5 == 0 {
            ans[i] = "Buzz"
        } else {
            ans[i] = strconv.Itoa(i + 1)
        }
    }

    return ans
}
