func mirrorReflection(p int, q int) int {
    _lcm := lcm(p, q)
    x, y := _lcm / q, _lcm / p;
    //fmt.Println(_lcm, x, y)
    if (y & 1) == 1 {
        if (x & 1) == 1 {
            return 1
        } else {
            return 2
        }
    }
    return 0
    
}

func lcm(a, b int) int {
    res := a * b
    for b != 0 {
        a, b = b, a % b
    }
    return res / a 
}
