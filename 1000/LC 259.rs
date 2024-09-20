impl Solution {
    pub fn three_sum_smaller(mut nums: Vec<i32>, target: i32) -> i32 {
        let n = nums.len();
        if n < 3 {
            return 0; 
        }
        let mut ans = 0;
        nums.sort_unstable();

        for i in 0..n-2 {
            if nums[i] * 3 >= target {
                break;
            }

            let (mut l, mut r) = (i + 1, n - 1);
            let x = target - nums[i];
            while l < r {
                if nums[l] + nums[r] < x {
                    ans += (r - l) as i32;
                    l += 1;
                } else {
                    r -= 1;
                }
            }
        }

        ans
    }
}