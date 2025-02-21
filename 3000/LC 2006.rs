const N: usize = 105;
impl Solution {
    pub fn count_k_difference(nums: Vec<i32>, k: i32) -> i32 {
        let mut cnt = vec![0; N];
        let mut ans = 0;
        for &x in &nums {
            if x - k >= 0 {
                ans += cnt[(x-k) as usize];
            } 
            if k + x < N as i32 {
                ans += cnt[(k+x) as usize];
            }
            cnt[x as usize] += 1;
        }

        ans
    }
}