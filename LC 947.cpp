using pii = pair<int,int>;
struct pair_hash {
    template <class T1, class T2>
    std::size_t operator () (const std::pair<T1,T2> &pair) const {
        auto hash1 = std::hash<T1>{}(pair.first);
        auto hash2 = std::hash<T2>{}(pair.second);
        return hash1 ^ hash2;
    }
};
class Union {
private:
    unordered_map<pii,pii,pair_hash> parent;
public:
    Union(){}
    pii find(pii x) {
        if (x != parent[x]) {
            parent[x] = find(parent[x]);
        }
        return parent[x];
    } 
    void merge(pii x, pii y) {
        pii rx = find(x);
        pii ry = find(y);
        if (rx != ry) {
            parent[ry] = rx; 
        }
    }
    void init(pii x) {
        parent[x] = x;
    }
};

class Solution {
public:
    int removeStones(vector<vector<int>>& stones) {
        int ans = 0;
        Union un;
        for (auto& st : stones) {
            un.init({st[0], st[1]});
        }
        sort(stones.begin(), stones.end(), [](const auto& a, const auto& b) {
            return a[0] < b[0];
        });
        for (size_t i = 0; i + 1 < stones.size(); i++) {
            if (stones[i][0] == stones[i+1][0]) {
                un.merge({stones[i][0],stones[i][1]}, {stones[i+1][0],stones[i+1][1]});
            }
        }
        sort(stones.begin(), stones.end(), [](const auto& a, const auto& b) {
            return a[1] < b[1];
        });
        for (size_t i = 0; i + 1 < stones.size(); i++) {
            if (stones[i][1] == stones[i+1][1]) {
                un.merge({stones[i][0],stones[i][1]}, {stones[i+1][0],stones[i+1][1]});
            }
        }
        unordered_map<pii, bool,pair_hash>m;
        for (auto& st : stones) {
            if (m[un.find({st[0], st[1]})] == false) {
                m[un.find({st[0], st[1]})] = true;
                ans++;
            } 
        }
        return stones.size() - ans;
    }
};