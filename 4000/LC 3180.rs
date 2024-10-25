impl Solution {
    pub fn max_total_reward(mut a: Vec<i32>) -> i32 {
        let mut f = vec![false; 4005];
        a.sort();
        f[0] = true;
        let mut ans = 0;

        for i in 0..a.len() {
            let mut x = a[i] - 1;
            while x >= 0 {
                if f[x as usize] {
                    f[(a[i] + x) as usize] = true;
                    ans = ans.max(a[i] + x);
                }
                x -= 1;
            }
        }

        ans
    }
}