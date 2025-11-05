const N int = 10_000
const M int = 1_000_000_000
func categorizeBox(length int, width int, height int, mass int) string {
    a, b := false, false
    if length >= N || width >= N || height >= N || length * width * height >= M {
        a = true
    } 
    if mass >= 100 {
        b = true
    } 

    if a && b {
        return "Both"
    } else if !a && !b {
        return "Neither"
    } else if !a && b {
        return "Heavy"
    }
    return "Bulky"
}