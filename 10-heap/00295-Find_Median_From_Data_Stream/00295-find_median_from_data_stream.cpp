// 295: Find Median From Data Stream
// https://leetcode.com/problems/find-median-from-data-stream

#include <queue>
#include <vector>
#include <iostream>

using namespace std;


// SOLUTION
class MedianFinder {
private:
    priority_queue<int> lower;
    priority_queue<int, vector<int>, greater<int>> higher;

public:
    MedianFinder(){}
    
    void addNum(int num) {
        if (lower.empty()) {
            lower.push(num);
            return;
        }
        
        if (lower.size() > higher.size()) {
            if (lower.top() > num) {
                higher.push(lower.top());
                lower.pop();
                lower.push(num);
            } else {
                higher.push(num);
            }
        } else {
            if (num > higher.top()) {
                lower.push(higher.top());
                higher.pop();
                higher.push(num);
            } else {
                lower.push(num);
            }
        }
    }
    
    double findMedian() {
        double result = 0.0;
        
        if (lower.size() == higher.size()) {
            result = lower.top() + (higher.top() - lower.top()) / 2.0;
        } else {
            if (lower.size() > higher.size()) {
                result = lower.top();
            } else {
                result = higher.top();
            }
        }
        
        return result;
    }
};


int main() {
    MedianFinder *o = new MedianFinder();

    // OPERATIONS
    o->addNum(1);
    o->addNum(2);
    cout << o->findMedian() << endl;
    o->addNum(3);
    cout << o->findMedian() << endl;

    return 0;
}