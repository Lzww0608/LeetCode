func maxBottlesDrunk(a int, b int) (ans int) {
    ans += a 

    for a >= b {
        ans++
        a = a - b + 1 
        b++
    }

    return 
}