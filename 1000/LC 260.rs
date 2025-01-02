impl Solution {
    pub fn single_number(nums: Vec<i32>) -> Vec<i32> {
        let mut xor = 0;
        for &x in &nums {
            xor ^= x;
        }

        let t = xor & (-xor);
        let mut a = 0;
        for &x in &nums {
            if x & t != 0 {
                a ^= x;
            }
        }

        [a, xor ^ a].to_vec()
    }
}