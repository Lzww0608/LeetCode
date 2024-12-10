const N: usize = 101;
impl Solution {
    pub fn sum_of_unique(nums: Vec<i32>) -> i32 {
        let mut cnt = vec![0; N];
        let mut ans = 0;
        for &x in &nums {
            let x = x as usize;
            cnt[x] += 1;
            if cnt[x] == 1 {
                ans += x;
            } else if cnt[x] == 2 {
                ans -= x;
            }
        }

        ans as _
    }
}