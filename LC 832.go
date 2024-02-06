func flipAndInvertImage(image [][]int) [][]int {
    n := len(image)
    for i := 0; i < n; i++ {
        for j := 0; j + j < n; j++ {
            image[i][j], image[i][n-j-1] = 1 ^ image[i][n-j-1], 1 ^ image[i][j]
        }
    }
    return image
}