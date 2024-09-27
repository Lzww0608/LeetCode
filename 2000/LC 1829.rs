impl Solution {
    pub fn get_maximum_xor(nums: Vec<i32>, maximum_bit: i32) -> Vec<i32> {
        let n = nums.len();
        let mut sum = 0;
        for &x in &nums {
            sum ^= x;
        }

        let mx = (1 << maximum_bit) - 1;
        let mut ans = vec![0; n];
        for i in 0..n {
            ans[i] = mx ^ sum;
            sum ^= nums[n-i-1];
        }

        ans
    }
}