impl Solution {
    pub fn difference_of_sum(nums: Vec<i32>) -> i32 {
        let mut ele_sum = 0;
        let mut dig_sum = 0;
        for &x in &nums {
            let mut x = x;
            ele_sum += x;
            while x > 0 {
                dig_sum += x % 10;
                x /= 10;
            }
        }

        ele_sum - dig_sum
    }
}