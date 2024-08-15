impl Solution {
    pub fn get_averages(nums: Vec<i32>, k: i32) -> Vec<i32> {
        let n = nums.len();
        let mut ans = vec![-1; n];
        let mut pre = vec![0i64; n + 1]; 

        for i in 0..n {
            pre[i + 1] = pre[i] + nums[i] as i64;  
        }

        let k = k as usize;
        if k > n {  
            return ans;
        }

        for i in k..n-k {
            let sum = pre[i + k + 1] - pre[i - k];  
            ans[i] = (sum / (2 * k as i64 + 1)) as i32; 
        }

        ans
    }
}
