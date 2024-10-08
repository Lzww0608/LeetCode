use std::collections::VecDeque;

impl Solution {
    pub fn max_score_indices(nums: Vec<i32>) -> Vec<i32> {
        let mut mx = 0;
        let mut pre = 0;
        let mut ans = VecDeque::new();
        ans.push_back(0);

        for (i, &x) in nums.iter().enumerate() {
            if x == 0 {
                pre += 1;
                if pre > mx {
                    mx = pre;
                    ans.clear();
                    ans.push_back(i as i32 + 1);
                } else if pre == mx {
                    ans.push_back(i as i32 + 1);
                }
            } else {
                pre -= 1;
            }
        }

        Vec::from(ans)
    }
}