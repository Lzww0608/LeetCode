class ExamRoom {
    int n;
    vector<int> a;
public:
    ExamRoom(int n) {
        this->n = n;
    }
    
    int seat() {
        if (a.empty()) {
            a.push_back(0);
            return 0;
        }
        int maxDist = a[0];
        int seatToSit = 0;
        for (int i = 0; i + 1 < a.size(); i++) {
            int dist = (a[i+1] - a[i]) / 2;
            if (dist > maxDist) {
                maxDist = dist;
                seatToSit = a[i] + dist;
            }
        }
        // check the back 
        if (n - 1 - a.back() > maxDist) {
            seatToSit = n - 1;
        }
        auto it = lower_bound(a.begin(), a.end(), seatToSit);
        a.insert(it, seatToSit);
        return seatToSit;
    }
    
    void leave(int p) {
        auto it = find(a.begin(), a.end(), p);
        a.erase(it);
    }
};

/**
 * Your ExamRoom object will be instantiated and called as such:
 * ExamRoom* obj = new ExamRoom(n);
 * int param_1 = obj->seat();
 * obj->leave(p);
 */