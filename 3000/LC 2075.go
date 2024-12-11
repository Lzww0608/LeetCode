
import "strings"
func decodeCiphertext(encodedText string, rows int) string {
    ans := []byte{}
    cols := len(encodedText) / rows
    for k := 0; k < cols; k++ {
        i, j := 0, k 
        for i * cols + j < len(encodedText) {
            ans = append(ans, encodedText[i*cols+j])
            i++
            j++
        }
    }

    return strings.TrimRight(string(ans), " ")
}