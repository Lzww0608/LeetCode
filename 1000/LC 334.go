
import "math"
func increasingTriplet(nums []int) bool {
    a, b := math.MaxInt, math.MaxInt 
    for _, x := range nums {
        if x > b {
            return true
        }

        if x < a {
            a = x 
        } else if x > a && x < b {
            b = x
        }
    }

    return false
}