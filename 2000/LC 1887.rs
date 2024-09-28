impl Solution {
    pub fn reduction_operations(mut nums: Vec<i32>) -> i32 {
        let n = nums.len();
        nums.sort_unstable();

        let mut ans = 0;
        let mut cnt = 0;
        for i in 1..n {
            if nums[i] != nums[i-1] {
                cnt += 1;
            }
            ans += cnt;
        }

        ans
    }
}