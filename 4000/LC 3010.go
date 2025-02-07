impl Solution {
    pub fn minimum_cost(nums: Vec<i32>) -> i32 {
        let n = nums.len();
        let mut ans = nums[0];
        let (mut x, mut y) = (i32::MAX, i32::MAX);
        for i in 1..n {
            if nums[i] < x {
                y = x; 
                x = nums[i];
            } else if nums[i] < y {
                y = nums[i];
            }
        }

        ans + x + y
    }
}