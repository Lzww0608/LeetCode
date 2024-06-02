class DinnerPlates {
    int capacity;
    vector<stack<int>> st;
    set<int> notEmpty;
    set<int> notFull;
public:
    DinnerPlates(int capacity) {
        this->capacity = capacity;

    }
    
    void push(int val) {
        if (notFull.size() == 0) {
            st.push_back(stack<int>());
            notFull.insert(st.size()-1);
        }
        auto t = *notFull.begin();
        st[t].push(val);
        if (st[t].size() == capacity) 
            notFull.erase(t);
        if (st[t].size() == 1)
            notEmpty.insert(t);
    }
    
    int pop() {
        if (notEmpty.size() == 0) 
            return -1;
        auto t = *notEmpty.rbegin();
        int res = st[t].top();
        st[t].pop();
        if (st[t].size() == 0) 
            notEmpty.erase(t);
        if (st[t].size() + 1 == capacity) 
            notFull.insert(t);
        return res;
    }
    
    int popAtStack(int index) {
        if (index >= st.size() || st[index].size() == 0)
            return -1;
        int res = st[index].top();
        st[index].pop();
        if (st[index].size() == 0)
            notEmpty.erase(index);
        if (st[index].size() + 1 == capacity)
            notFull.insert(index);
        return res;
    }
};

/**
 * Your DinnerPlates object will be instantiated and called as such:
 * DinnerPlates* obj = new DinnerPlates(capacity);
 * obj->push(val);
 * int param_2 = obj->pop();
 * int param_3 = obj->popAtStack(index);
 */
