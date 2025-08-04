func numOfUnplacedFruits(fruits []int, baskets []int) (ans int) {
    n := len(fruits)
    for _, x := range fruits {
        f := false
        for i := 0; i < n; i++ {
            if baskets[i] >= x {
                baskets[i] = 0
                f = true 
                break
            }
        }

        if !f {
            ans++
        }
    }

    return
}