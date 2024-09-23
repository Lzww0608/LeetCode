impl Solution {
    pub fn count_subarrays(nums: Vec<i32>, min_k: i32, max_k: i32) -> i64 {
        let mut l = -1;
        let (mut mn, mut mx) = (-1, -1);
        let mut ans: i64 = 0;
        
        for (i, &x) in nums.iter().enumerate() {
            if x < min_k || x > max_k {
                l = i as i64;
                mn = i as i64;
                mx = i as i64;
            } else {
                if x == max_k {
                    mx = i as i64;
                }
                if x == min_k {
                    mn = i as i64;
                }
                ans += mn.min(mx) as i64 - l as i64;

            }

        }

        ans
    }
}