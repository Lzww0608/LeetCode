class MedianFinder {
    priority_queue<int> pq1;
    priority_queue<int, vector<int>, greater<int>> pq2;
public:
    MedianFinder() {

    }
    
    void addNum(int num) {
        if (pq1.empty() || num < pq1.top() ) {
            pq1.push(num);
        } else {
            pq2.push(num);
        }
        while (pq1.size() > pq2.size() + 1) {
            auto t = pq1.top();
            pq1.pop();
            pq2.push(t);
        }
        while (pq1.size() < pq2.size()) {
            auto t = pq2.top();
            pq2.pop();
            pq1.push(t);
        }
    }
    
    double findMedian() {
        if (pq1.size() == pq2.size()) {
            return 1.0 * (pq1.top() + pq2.top()) / 2;
        } 
        return pq1.top();
    }
};

/**
 * Your MedianFinder object will be instantiated and called as such:
 * MedianFinder* obj = new MedianFinder();
 * obj->addNum(num);
 * double param_2 = obj->findMedian();
 */