impl Solution {
    pub fn sort_transformed_array(nums: Vec<i32>, a: i32, b: i32, c: i32) -> Vec<i32> {
        let n = nums.len();
        let (mut l, mut r) = (0, n - 1);
        let mut ans = vec![0; n];

        let mut k = if a >= 0 {
            r
        } else {
            l 
        };

        let f = |x: i32| -> i32 {
            a * x * x + b * x + c
        };

        while l <= r {
            if a >= 0 {
                let x = f(nums[l]);
                let y = f(nums[r]);
                if x >= y {
                    ans[k] = x;
                    l += 1;
                } else {
                    ans[k] = y;
                    r -= 1;
                }
                k -= 1
            } else {
                let x = f(nums[l]);
                let y = f(nums[r]);
                if x <= y {
                    ans[k] = x;
                    l += 1;
                } else {
                    ans[k] = y;
                    r -= 1;
                }
                k += 1
            }
        }

        ans
    }
}