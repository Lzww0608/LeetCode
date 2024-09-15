impl Solution {
    pub fn number_of_points(nums: Vec<Vec<i32>>) -> i32 {
        let mut ans = 0;
        let mut cnt = vec![0; 105];
        for v in &nums {
            cnt[v[0] as usize] += 1;
            cnt[v[1] as usize + 1] -= 1;
        }

        let mut cur = 0;
        for &x in &cnt {
            cur += x;
            if cur > 0 {
                ans += 1;
            }
        }

        ans
    }
}