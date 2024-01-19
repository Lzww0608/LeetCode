class TopVotedCandidate {
    vector<int> times;
    vector<int> res;
public:
    TopVotedCandidate(vector<int>& persons, vector<int>& times) {
        int n = persons.size();
        vector<int> ans(n, 0);
        vector<int> res(n, 0);
        int cnt = 0, p = 0;
        for (int i = 0; i < n; ++i) {
            ans[persons[i]]++;
            if (cnt <= ans[persons[i]]) {
                cnt = ans[persons[i]];
                p = persons[i];
            } else if (p == persons[i]) {
                cnt = persons[i];
            }
            res[i] = p;
        }
        this->times = times;
        this->res = res;
    }
    
    int q(int t) {
        int l = 0, r = times.size() - 1;
        while (l < r) {
            int mid = l + ((r - l + 1) >> 1);
            if (times[mid] > t) {
                r = mid - 1;
            } else if (times[mid] < t) {
                l = mid;
            } else {
                return res[mid];
            }
        }
        return res[l];
        /*int pos = upper_bound(times.begin(), times.end(), t) - times.begin() - 1;
        return res[pos];*/
    }
};

/**
 * Your TopVotedCandidate object will be instantiated and called as such:
 * TopVotedCandidate* obj = new TopVotedCandidate(persons, times);
 * int param_1 = obj->q(t);
 */