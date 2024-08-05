impl Solution {
    pub fn bag_of_tokens_score(tokens: Vec<i32>, power: i32) -> i32 {
        if tokens.len() == 0 {
            return 0;
        }
        let mut power = power;
        let mut tokens = tokens;
        tokens.sort();
        let (mut l, mut r) = (0, tokens.len() - 1);
        let (mut cur, mut ans) = (0, 0);

        while l <= r {
            if power >= tokens[l] {
                power -= tokens[l];
                l += 1;
                cur += 1;
                ans = ans.max(cur);
            } else if cur > 0 {
                cur -= 1;
                power += tokens[r];
                r -= 1;
            } else {
                break;
            }
        }

        ans
    }
}