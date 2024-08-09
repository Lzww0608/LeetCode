impl Solution {
    pub fn broken_calc(mut start_value: i32, mut target: i32) -> i32 {
        let mut ans: i32 = 0;
        while target > start_value {
            if (target & 1) == 0 {
                target >>= 1;
                ans += 1;
            } else {
                target = (target + 1) >> 1;
                ans += 2;
            }
        }

        if start_value > target {
            ans += start_value - target
        }

        ans
    }
}