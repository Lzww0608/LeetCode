impl Solution {
    pub fn triangular_sum(mut nums: Vec<i32>) -> i32 {
        let n = nums.len();
        for i in 0..n-1 {
            for j in 0..n-i-1 {
                nums[j] += nums[j+1];
                nums[j] %= 10;
            }
        }

        nums[0]
    }
}