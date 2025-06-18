impl Solution {
    pub fn divide_array(mut nums: Vec<i32>, k: i32) -> Vec<Vec<i32>> {
        nums.sort_unstable();
        let n = nums.len();
        let mut ans = Vec::new();
        for i in (2..n).step_by(3) {
            if nums[i] - nums[i-2] <= k {
                ans.push(vec![nums[i-2], nums[i-1], nums[i]]);
            } else {
                return Vec::new();
            }
        }

        ans
    }
}