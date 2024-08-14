impl Solution {
    pub fn find_different_binary_string(nums: Vec<String>) -> String {
        let mut ans = vec!['0'; nums.len()];

        for (i, x) in nums.iter().enumerate() {
            if x.chars().nth(i).unwrap() == '0' {
                ans[i] = '1';
            } else {
                ans[i] = '0';
            }
        }

        ans.into_iter().collect()
    }
}