func minOperations(nums []int, numsDivide []int) int {
    sort.Ints(nums)
    x := 0
    for _, t := range numsDivide {
        x = gcd(t, x)
    }


    for i, t := range nums {
        if x % t == 0 {
            return i
        } 
    }

    return -1
}

func gcd(x, y int) int {
    for y > 0 {
        x, y = y, x % y
    }

    return x
}