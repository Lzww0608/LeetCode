impl Solution {
    pub fn find_array(pref: Vec<i32>) -> Vec<i32> {
        let n = pref.len();
        let mut ans = vec![pref[0]; n];

        for i in 1..n {
            ans[i] = pref[i] ^ pref[i-1]
        }

        ans
    }
}