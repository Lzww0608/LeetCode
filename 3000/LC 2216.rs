impl Solution {
    pub fn min_deletion(nums: Vec<i32>) -> i32 {
        let mut ans = 0;
        let n = nums.len();
        for i in 0..n {
            if (i - ans as usize) & 1 == 0 && i + 1 < n && nums[i] == nums[i+1] {
                ans += 1;
            }
        }
        if (n - ans as usize) & 1 == 1 {
            ans += 1;
        }
        ans
    }
}