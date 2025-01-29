/** 
 * Definition of commonBits API.
 * func commonBits(num int) int;
 */
const N int = 30
func findNumber() (ans int) {
    for i := 0; i <= N; i++ {
        a := commonBits(0)
        b := commonBits(1 << i)
        if b - a == 1 {
            ans |= 1 << i 
        } 
    }

    return
}