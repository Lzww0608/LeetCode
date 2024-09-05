impl Solution {
    pub fn min_meeting_rooms(intervals: Vec<Vec<i32>>) -> i32 {
        let mut ans: i32 = 0;
        let mut cur: i32 = 0;
        let mut mx: i32 = 0;

        for v in &intervals {
            mx = mx.max(v[1]);
        }

        let mut m = vec![0; mx as usize + 1];
        for v in &intervals {
            m[v[0] as usize] += 1;
            m[v[1] as usize] -= 1;
        }

        for x in &m {
            cur += x;
            ans = ans.max(cur);
        }

        ans
    }
}