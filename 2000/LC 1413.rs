impl Solution {
    pub fn min_start_value(nums: Vec<i32>) -> i32 {
        let mut ans = 1;
        let mut cur = 0;
        for &x in &nums {
            cur += x;
            ans = ans.max(-cur + 1);
        }

        ans
    }
}