impl Solution {
    pub fn occurrences_of_element(nums: Vec<i32>, queries: Vec<i32>, x: i32) -> Vec<i32> {
        let n = queries.len();
        let mut ans = vec![-1; n];
        let mut pos = Vec::new();

        for (i, &t) in nums.iter().enumerate() {
            if t == x {
                pos.push(i);
            }
        }

        for i in 0..queries.len() {
            let t = queries[i];
            if t <= pos.len() as i32 {
                ans[i] = pos[(queries[i] - 1) as usize] as i32;
            } 
        }

        ans
    }
}