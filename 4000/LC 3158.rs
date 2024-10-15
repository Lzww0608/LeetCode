impl Solution {
    pub fn duplicate_numbers_xor(nums: Vec<i32>) -> i32 {
        let mut ans = 0i32;
        let mut vis = 0i64;

        for &x in &nums {
            if (vis >> x) & 1 == 1 {
                ans ^= x;
            } else {
                vis |= 1 << x;
            }
        }

        ans
    }
}