impl Solution {
    pub fn min_size_subarray(nums: Vec<i32>, mut target: i32) -> i32 {
        let sum: i64 = nums.iter().map(|&x| x as i64).sum();
        let n = nums.len();

        let mut cur = 0;
        let mut mn = n;
        let mut l = 0;
        let rem = (target as i64 % sum) as i32;
        for r in 0..n * 2 {
            let x = nums[r % n];
            cur += x;
            while cur > rem {
                cur -= nums[l % n];
                l += 1;
            }
            if cur == rem {
                mn = mn.min(r - l + 1);
            }
        }

        if mn >= n {
            return -1;
        }

        mn as i32 + (target as i64 / sum) as i32 * n as i32
    }
}