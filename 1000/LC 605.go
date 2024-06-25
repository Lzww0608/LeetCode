func canPlaceFlowers(flowerbed []int, n int) bool {
    m := len(flowerbed)
    for i := range flowerbed {
        if (i == 0 || flowerbed[i-1] == 0) && flowerbed[i] == 0 && (i == m - 1 || flowerbed[i+1] == 0) {
            flowerbed[i] = 1
            n--
            if n <= 0 {
                return true
            }
        }
    } 

    return n <= 0
}
