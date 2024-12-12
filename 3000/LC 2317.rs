impl Solution {
    pub fn maximum_xor(nums: Vec<i32>) -> i32 {
        let mut ans = 0;
        for &x in &nums {
            ans |= x;
        }

        ans
    }
}