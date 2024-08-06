use std::collections::HashMap;
impl Solution {
    pub fn get_distances(arr: Vec<i32>) -> Vec<i64> {
        let mut m: HashMap<i32, Vec<usize>> = HashMap::new();

        for (i, &x) in arr.iter().enumerate() {
            m.entry(x).or_insert(Vec::new()).push(i);
        }

        let mut ans: Vec<i64> = vec![0; arr.len()];
        for v in m.values() {
            let mut sum: i64 = 0;
            for &i in v.iter() {
                sum += i as i64 - v[0] as i64
            }
            ans[v[0]] = sum;
            for i in 1..v.len() {
                sum += (i * 2 - v.len()) as i64 * (v[i] - v[i-1]) as i64;
                ans[v[i]] = sum;
            }
        }

        ans
    }
}