impl Solution {
    pub fn create_target_array(nums: Vec<i32>, index: Vec<i32>) -> Vec<i32> {
        let n = nums.len();
        let mut ans = vec![0; n];
        for i in 0..n {
            let id = index[i] as usize;
            let x = nums[i];
            for j in (id+1..n).rev() {
                ans[j] = ans[j-1];
            }
            ans[id] = x;
        }

        ans
    }
}