impl Solution {
    pub fn wiggle_sort(nums: &mut Vec<i32>) {
        let n = nums.len();

        for i in 1..n {
            if i % 2 == 1 && nums[i] < nums[i-1] || i % 2 == 0 && nums[i] > nums[i-1] {
                nums.swap(i, i - 1);
            }
        }

        
    }
}