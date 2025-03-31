func smallestFactorization(num int) int {
    if num < 10 {
        return num
    }

    a := []int{}
    for i := 9; i > 1; i-- {
        cnt := 0
        for num % i == 0 {
            num /= i 
            cnt++
            
        }
        for j := 0; j < cnt; j += 1 {
            a = append(a, i)
        }
    }

    if num > 1 {
        return 0
    }

    ans := 0
    for i := len(a) - 1; i >= 0; i-- {
        ans = ans * 10 + a[i]
        if ans > math.MaxInt32 {
            return 0
        }
    }

    return ans
}