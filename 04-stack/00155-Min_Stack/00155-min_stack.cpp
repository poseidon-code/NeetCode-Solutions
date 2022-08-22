// 155: Min stack
// https://leetcode.com/problems/min-stack/


#include<iostream>
#include<stack>

using namespace std;

// SOLUTION
class MinStack {
    private:
        stack<int> s1, s2;
    
    public:
        void push(int x) {
            s1.push(x);
            if (s2.empty() || x <= getMin())
                s2.push(x);
        }

        void pop() {
            if (s1.top() == getMin())
                s2.pop();
            s1.pop();
        }

        int top() {
            return s1.top();
        }

        int getMin() {
            return s2.top();
        }
};


int main() {
    MinStack *o = new MinStack();

    // OPERATIONS
    o->push(-2);
    o->push(0);
    o->push(-3);
    cout<<o->getMin()<<" ";
    o->pop();
    o->top();
    cout<<o->getMin()<<" ";

    cout<<endl;

    return 0;
}