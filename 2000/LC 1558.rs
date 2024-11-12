impl Solution {
    pub fn min_operations(nums: Vec<i32>) -> i32 {
        let mut ans = 0;
        let mut mx = 0;

        for &x in &nums {
            let mut x = x;
            let mut cur = 0;
            while x > 0 {
                if x % 2 == 1 {
                    x -= 1;
                    ans += 1;
                } else {
                    x >>= 1;
                    cur += 1;
                }
            }   
            mx = mx.max(cur);
        }

        ans + mx
    }
}