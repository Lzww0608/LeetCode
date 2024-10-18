impl Solution {
    pub fn min_operations(nums: Vec<i32>) -> i32 {
        let mut ans = 0;
        for &x in &nums {
            if x == ans & 1 {
                ans += 1;
            }
        }

        ans
    }
}