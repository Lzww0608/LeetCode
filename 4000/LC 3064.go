/** 
 * Definition of commonSetBits API.
 * func commonSetBits(num int) int;
 */

const N int = 30

func findNumber() (ans int) {
    for i := 0; i <= N; i++ {
        if commonSetBits(1 << i) == 1 {
            ans |= 1 << i
        }
    }

    return
}