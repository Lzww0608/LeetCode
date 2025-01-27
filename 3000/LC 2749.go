impl Solution {
    pub fn make_the_integer_zero(num1: i32, num2: i32) -> i32 {
        let num1 = num1 as i64;
        let num2 = num2 as i64;
        let mut k = 1;
        while k <= num1 - num2 * k {
            if k >= (num1 - num2 * k).count_ones() as i64 {
                return k as _;
            }
            k += 1;
        }

        -1
    }
}