func maxIceCream(costs []int, coins int) (ans int) {
    sort.Ints(costs)
    for _, x := range costs {
        if x > coins {
            break
        }
        coins -= x 
        ans++
    }

    return
}