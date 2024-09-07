impl Solution {
    pub fn find_max_consecutive_ones(nums: Vec<i32>) -> i32 {
        let mut pre = -1;
        let mut cur = 0;
        let mut ans = 0;

        for x in nums {
            if x == 0 {
                pre = cur;
                cur = 0;
            } else {
                cur += 1;
            }

            ans = ans.max(cur + pre + 1);
        }

        ans
    }
}