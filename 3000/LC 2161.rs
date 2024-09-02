impl Solution {
    pub fn pivot_array(nums: Vec<i32>, pivot: i32) -> Vec<i32> {
        let mut ans = Vec::new();
        ans.reserve(nums.len());
        for &x in &nums {
            if x < pivot {
                ans.push(x);
            }
        }

        for &x in &nums {
            if x == pivot {
                ans.push(x);
            }
        }

        for &x in &nums {
            if x > pivot {
                ans.push(x);
            }
        }

        ans
    }
}