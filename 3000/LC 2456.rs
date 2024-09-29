use std::collections::HashMap;

impl Solution {
    pub fn most_popular_creator(creators: Vec<String>, ids: Vec<String>, views: Vec<i32>) -> Vec<Vec<String>> {
        let n = creators.len();
        let mut total_views: HashMap<String, i32> = HashMap::new(); 
        let mut c2i: HashMap<String, String> = HashMap::new();     
        let mut cnt: HashMap<String, i32> = HashMap::new();         
        let mut max_views = 0;

        for i in 0..n {
            let creator = &creators[i];
            let id = &ids[i];
            let view = views[i];

            *total_views.entry(creator.clone()).or_insert(0) += view;
            max_views = max_views.max(total_views[creator]);

            if !c2i.contains_key(creator) || view > *cnt.entry(creator.clone()).or_insert(0) || (view == cnt[creator] && id.as_bytes() < &c2i[creator].as_bytes()) {
                c2i.insert(creator.clone(), id.clone());
            }
            cnt.insert(creator.clone(), cnt.get(creator).cloned().unwrap_or(0).max(view));
        }

        let mut ans = Vec::new();
        for (creator, &view) in &total_views {
            if view == max_views {
                ans.push(vec![creator.clone(), c2i[creator].clone()]);
            }
        }

        if ans.is_empty() {
            ans.push(vec!["a".to_string(), "a".to_string()]);
        }

        ans
    }
}