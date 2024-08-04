impl Solution {
    pub fn max_product(nums: Vec<i32>) -> i32 {
        let mut a: i32 = 0;
        let mut b: i32 = 0;
        for &x in nums.iter() {
            if x > a {
                b = a;
                a = x;
            } else if x > b {
                b = x
            }
        }

        (a - 1) * (b - 1)
    }
}