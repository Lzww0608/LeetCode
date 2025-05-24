impl Solution {
    pub fn largest_subarray(nums: Vec<i32>, k: i32) -> Vec<i32> {
        let mut n = nums.len();
        let mut start = 0;
        for i in 0..(n - k as usize + 1) {
            if nums[start] < nums[i] {
                start = i;
            }
        }

        nums[start..start + k as usize].to_vec()
    }
}impl Solution {
    pub fn largest_subarray(nums: Vec<i32>, k: i32) -> Vec<i32> {
        let mut n = nums.len();
        let mut start = 0;
        for i in 0..(n - k as usize + 1) {
            if nums[start] < nums[i] {
                start = i;
            }
        }

        nums[start..start + k as usize].to_vec()
    }
}