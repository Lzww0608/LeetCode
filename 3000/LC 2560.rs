impl Solution {
    pub fn min_capability(nums: Vec<i32>, k: i32) -> i32 {
        let mut l: i32 = 1;
        let mut r: i32 = 1_000_000_000;

        fn check(target: i32, nums: &Vec<i32>, k: i32) -> bool {
            let mut cnt = 0;
            let mut pre = -2;
            for (i, &x) in nums.iter().enumerate() {
                if x <= target && i as i32 > pre + 1 {
                    pre = i as i32;
                    cnt += 1;
                }
            }

            cnt >= k
        }

        while l < r {
            let mid = l + ((r - l) >> 1);
            if check(mid, &nums, k) {
                r = mid 
            } else {
                l = mid + 1;
            }
        }

        l
    }
}